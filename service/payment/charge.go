package payment

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type CreateChargeInput struct {
	Type      string             `json:"type"`
	Mode      string             `json:"mode"`
	Payer     *ChargePayerInput  `json:"payer"`
	Device    *ChargeDeviceInput `json:"device"`
	Order     *ChargeOrderInput  `json:"order"`
	Amount    *ChargeAmountInput `json:"amount"`
	NotifyUrl string             `json:"notify_url"`
}

type ChargePayerInput struct {
	OpenID   string `json:"open_id"`
	ClientIP string `json:"client_ip"`
}

type ChargeDeviceInput struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type ChargeOrderInput struct {
	Number  string `json:"number"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ChargeAmountInput struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

type CreateChargeOutput struct {
	Type        string                   `json:"type"`
	Mode        string                   `json:"mode"`
	Transaction *ChargeTransactionOutput `json:"transaction"`
	Order       *ChargeOrderOutput       `json:"order"`
	Amount      *ChargeAmountOutput      `json:"amount"`
	Credential  any                      `json:"credential"`
}

type ChargeTransactionOutput struct {
	Number string `json:"number"`
}

type ChargeOrderOutput struct {
	Number  string `json:"number"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ChargeAmountOutput struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

func (c *Client) CreateCharge(ctx context.Context, input *CreateChargeInput) (*CreateChargeOutput, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodPost).
		SetPath("/charges").
		SetBody(body)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	output := new(CreateChargeOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
