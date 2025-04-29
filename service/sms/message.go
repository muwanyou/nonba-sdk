package sms

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type SendMessageParams struct {
	TemplateID     int64    `json:"template_id,omitempty"`
	SignatureID    int64    `json:"signature_id,omitempty"`
	PhoneNumber    string   `json:"phone_number,omitempty"`
	TemplateParams []string `json:"template_params,omitempty"`
}

type SendMessageResult struct {
	ID int64 `json:"id,string,omitempty"`
}

func (c *Client) SendMessage(ctx context.Context, params *SendMessageParams) (*SendMessageResult, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodPost).
		SetPath("/messages").
		SetBody(body)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(SendMessageResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
