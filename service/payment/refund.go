package payment

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/muwanyou/nonba-sdk/enum"
)

type CreateRefundParam struct {
	Reason      string                  `json:"reason"`
	ClientIP    string                  `json:"client_ip"`
	Transaction *RefundTransactionParam `json:"transaction"`
	Order       *RefundOrderParam       `json:"order"`
	Amount      *RefundAmountParam      `json:"amount"`
	NotifyUrl   string                  `json:"notify_url"`
}

type RefundTransactionParam struct {
	Number string `json:"number"`
}

type RefundOrderParam struct {
	Number string `json:"number"`
}

type RefundAmountParam struct {
	Currency string `json:"currency"`
	Total    int64  `json:"total"`
}

type CreateRefundResult struct {
	Transaction *RefundTransactionParam `json:"transaction"`
	Order       *RefundOrderParam       `json:"order"`
}

type RefundTransactionResult struct {
	Number string `json:"number"`
}

type RefundOrderResult struct {
	Number string `json:"number"`
}

func (c *Client) CreateRefund(ctx context.Context, param *CreateRefundParam) (*CreateRefundResult, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodPost).
		SetPath(fmt.Sprintf("/charges/%s/refunds", param.Transaction.Number)).
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
