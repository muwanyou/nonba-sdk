package algorithm

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
		// WithHost("api.algorithm.nonba.net").
		// WithPort(443).
		WithScheme(enum.SchemeHttp).
		WithHost("120.36.140.210").
		WithPort(8010).
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
