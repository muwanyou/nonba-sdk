package sms

import (
	"github.com/muwanyou/nonba-sdk/core"
	"github.com/muwanyou/nonba-sdk/enum"
)

type Client struct {
	core.Client
}

func NewClient(credential *core.Credential) *Client {
	client := new(Client)
	client.Init().WithCredential(credential)
	return client
}

type Request struct {
	core.BaseRequest
}

func NewRequest() *Request {
	request := new(Request)
	request.Init().
		// WithScheme(enum.SchemeHttps).
		// WithHost("api.sms.nonba.net").
		// WithPort(443).
		WithScheme(enum.SchemeHttp).
		WithHost("112.5.167.217").
		WithPort(8040).
		WithVersion("1")
	return request
}

type Response struct {
	core.BaseResponse
}

func NewResponse() *Response {
	response := new(Response)
	return response
}
