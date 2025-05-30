package carrier

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type SubmitVerificationInput struct {
	FamilyName         string `json:"family_name"`
	GivenName          string `json:"given_name"`
	IdentityCardNumber string `json:"identity_card_number"`
	PhoneNumber        string `json:"phone_number"`
}

type SubmitVerificationOutput struct {
}

func (c *Client) SubmitVerification(ctx context.Context, input *SubmitVerificationInput) (*SubmitVerificationOutput, error) {
	body, err := json.Marshal(input)
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
	output := new(SubmitVerificationOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
