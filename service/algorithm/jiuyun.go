package algorithm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

type GetJiuyunInput struct {
	FamilyName string    `json:"family_name"`
	GivenName  string    `json:"given_name"`
	Sex        enum.Sex  `json:"sex"`
	Birthday   time.Time `json:"birthday"`
	Datetime   time.Time `json:"datetime"`
}

type GetJiuyunOutput struct {
	Impact   string `json:"impact"`
	Epoch    string `json:"epoch"`
	PassYear int8   `json:"pass_year"`
}

func (c *Client) GetJiuyun(ctx context.Context, input *GetJiuyunInput) (*GetJiuyunOutput, error) {
	bytes, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	query := make(map[string]string)
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		return nil, err
	}
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodGet).
		SetPath("/jiuyun").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	output := new(GetJiuyunOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
