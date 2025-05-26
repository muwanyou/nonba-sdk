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
		// 120.36.140.210
		WithScheme(enum.SchemeHttp).
		WithHost("127.0.0.1").
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

type Dimension struct {
	ID      int64   `json:"id,string"`
	Name    string  `json:"name"`
	Alias   string  `json:"alias"`
	Entries []Entry `json:"entries"`
}

type Entry struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
