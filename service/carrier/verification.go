package carrier

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type SubmitVerificationParam struct {
	PhoneNumber        string `json:"phone_number"`
	FullName           string `json:"full_name"`
	IdentityCardNumber string `json:"identity_card_number"`
}

type SubmitVerificationResult struct {
}

func (c *Client) SubmitVerification(ctx context.Context, params *SubmitVerificationParam) (*SubmitVerificationResult, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodPost).
		SetPath("/verifications").
		SetBody(body)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(SubmitVerificationResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
