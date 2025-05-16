package payment

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type CreateChargeParam struct {
	Type      string             `json:"type"`
	Mode      string             `json:"mode"`
	Payer     *ChargePayerParam  `json:"payer"`
	Device    *ChargeDeviceParam `json:"device"`
	Order     *ChargeOrderParam  `json:"order"`
	Amount    *ChargeAmountParam `json:"amount"`
	NotifyUrl string             `json:"notify_url"`
}

type ChargePayerParam struct {
	OpenID   string `json:"open_id"`
	ClientIP string `json:"client_ip"`
}

type ChargeDeviceParam struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

type ChargeOrderParam struct {
	Number  string `json:"number"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ChargeAmountParam struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

type CreateChargeResult struct {
	Type        string                   `json:"type"`
	Mode        string                   `json:"mode"`
	Transaction *ChargeTransactionResult `json:"transaction"`
	Order       *ChargeOrderResult       `json:"order"`
	Amount      *ChargeAmountResult      `json:"amount"`
	Credential  any                      `json:"credential"`
}

type ChargeTransactionResult struct {
	Number string `json:"number"`
}

type ChargeOrderResult struct {
	Number  string `json:"number"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type ChargeAmountResult struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

func (c *Client) CreateCharge(ctx context.Context, param *CreateChargeParam) (*CreateChargeResult, error) {
	body, err := json.Marshal(param)
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
	result := new(CreateChargeResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
