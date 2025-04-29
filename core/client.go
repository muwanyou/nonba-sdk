package core

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/muwanyou/nonba-sdk/constant"
	"github.com/muwanyou/nonba-sdk/enum"
)

type Client struct {
	http   *http.Client
	signer Signer
}

func (c *Client) Init() *Client {
	if c.http == nil {
		c.http = http.DefaultClient
	}
	return c
}

func (c *Client) WithSigner(signer Signer) *Client {
	c.signer = signer
	return c
}

func (c *Client) WithCredential(credential *Credential) *Client {
	signer := NewHMACSigner(credential)
	return c.WithSigner(signer)
}

func (c *Client) Send(request Request, response Response) error {
	httpRequest, err := http.NewRequestWithContext(request.GetContext(), request.GetMethod().String(), request.GetUrl(), bytes.NewBuffer(request.GetBody()))
	if err != nil {
		return err
	}
	for key, value := range request.GetHeader() {
		httpRequest.Header.Set(key, value)
	}
	if request.GetMethod() == enum.MethodGet {
		httpRequest.Header.Set("Content-Type", constant.ContentTypeForm)
	} else {
		httpRequest.Header.Set("Content-Type", constant.ContentTypeJSON)
	}
	httpRequest.Header.Set("Authorization", fmt.Sprintf("%s %s", c.signer.Type(), c.signer.Sign()))
	httpResponse, err := c.http.Do(httpRequest)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return err
	}
	response.SetCode(int16(httpResponse.StatusCode))
	response.SetBody(body)
	if response.GetCode() != http.StatusOK {
		return response.GetError()
	}
	return nil
}
