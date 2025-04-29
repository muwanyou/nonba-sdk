package payment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/muwanyou/nonba-sdk/enum"
)

type CreateRefundParams struct {
	Reason      string                   `json:"reason"`
	ClientIP    string                   `json:"client_ip"`
	Transaction *RefundTransactionParams `json:"transaction"`
	Order       *RefundOrderParams       `json:"order"`
	Amount      *RefundAmountParams      `json:"amount"`
	NotifyUrl   string                   `json:"notify_url"`
}

type RefundTransactionParams struct {
	Number string `json:"number"`
}

type RefundOrderParams struct {
	Number string `json:"number"`
}

type RefundAmountParams struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

type CreateRefundResult struct {
	Transaction *RefundTransactionParams `json:"transaction"`
	Order       *RefundOrderParams       `json:"order"`
}

type RefundTransactionResult struct {
	Number string `json:"number"`
}

type RefundOrderResult struct {
	Number string `json:"number"`
}

func (c *Client) CreateRefund(ctx context.Context, params *CreateRefundParams) (*CreateRefundResult, error) {
	body, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodPost).
		SetPath(fmt.Sprintf("/charges/%s/refunds", params.Transaction.Number)).
		SetBody(body)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(CreateRefundResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
