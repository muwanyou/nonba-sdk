package core

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response interface {
	SetCode(int16) *BaseResponse
	SetBody([]byte) *BaseResponse
	GetCode() int16
	GetBody() []byte
	GetError() error
}

type ErrorResult struct {
	Code     int16  `json:"code"`
	Reason   string `json:"reason"`
	Message  string `json:"message"`
	Metadata any    `json:"metadata"`
}

type BaseResponse struct {
	code int16
	body []byte
}

func NewBaseResponse() *BaseResponse {
	return new(BaseResponse)
}

func (r *BaseResponse) SetCode(code int16) *BaseResponse {
	r.code = code
	return r
}

func (r *BaseResponse) SetBody(body []byte) *BaseResponse {
	r.body = body
	return r
}

func (r *BaseResponse) GetCode() int16 {
	return r.code
}

func (r *BaseResponse) GetBody() []byte {
	return r.body
}

func (r *BaseResponse) GetError() error {
	if r.GetCode() == http.StatusOK {
		return nil
	}
	result := new(ErrorResult)
	err := json.Unmarshal(r.GetBody(), &result)
	if err != nil {
		return err
	}
	return fmt.Errorf("code=%d,reason=%s,message=%s", result.Code, result.Reason, result.Message)
}
