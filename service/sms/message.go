package sms

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type SendMessageInput struct {
	TemplateID     int64    `json:"template_id,omitempty"`
	SignatureID    int64    `json:"signature_id,omitempty"`
	PhoneNumber    string   `json:"phone_number,omitempty"`
	TemplateInputs []string `json:"template_Inputs,omitempty"`
}

type SendMessageOutput struct {
	ID int64 `json:"id,string,omitempty"`
}

func (c *Client) SendMessage(ctx context.Context, input *SendMessageInput) (*SendMessageOutput, error) {
	body, err := json.Marshal(input)
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
	output := new(SendMessageOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
