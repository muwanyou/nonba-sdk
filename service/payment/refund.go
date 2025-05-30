package payment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/muwanyou/nonba-sdk/enum"
)

type CreateRefundInput struct {
	Reason      string                  `json:"reason"`
	ClientIP    string                  `json:"client_ip"`
	Transaction *RefundTransactionInput `json:"transaction"`
	Order       *RefundOrderInput       `json:"order"`
	Amount      *RefundAmountInput      `json:"amount"`
	NotifyUrl   string                  `json:"notify_url"`
}

type RefundTransactionInput struct {
	Number string `json:"number"`
}

type RefundOrderInput struct {
	Number string `json:"number"`
}

type RefundAmountInput struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

type CreateRefundOutput struct {
	Transaction *RefundTransactionInput `json:"transaction"`
	Order       *RefundOrderInput       `json:"order"`
}

type RefundTransactionOutput struct {
	Number string `json:"number"`
}

type RefundOrderOutput struct {
	Number string `json:"number"`
}

func (c *Client) CreateRefund(ctx context.Context, input *CreateRefundInput) (*CreateRefundOutput, error) {
	body, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodPost).
		SetPath(fmt.Sprintf("/charges/%s/refunds", input.Transaction.Number)).
		SetBody(body)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	output := new(CreateRefundOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
