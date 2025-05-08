package payment

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
		WithScheme(enum.SchemeHttps).
		WithHost("api.payment.nonba.net").
		WithPort(443).
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
