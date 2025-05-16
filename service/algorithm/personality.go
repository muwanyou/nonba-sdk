package algorithm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

type Personalitiey struct {
	ID          int64
	Name        string
	Alias       string
	Positive    string
	Negative    string
	Solve       string
	Description string
}

// 性格列表参数
type ListPersonalitiesParams struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
}

// 性格列表结果
type ListPersonalitiesResult struct {
	Items []*Personalitiey
}

// 获取性格列表
func (c *Client) ListPersonalities(ctx context.Context, params *ListPersonalitiesParams) (*ListPersonalitiesResult, error) {
	bytes, err := json.Marshal(params)
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
		SetPath("/personalities").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(ListPersonalitiesResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 性格纬度列表参数
type ListPersonalityDimensionsParams struct {
	FamilyName string    `json:"family_name"`
	GivenName  string    `json:"given_name"`
	Sex        enum.Sex  `json:"sex"`
	Birthday   time.Time `json:"birthday"`
	Datetime   time.Time `json:"datetime"`
	Height     int16     `json:"height"`
	Weight     int16     `json:"weight"`
}

// 性格纬度列表结果
type ListPersonalityDimensionsResult struct {
	Before []Dimension `json:"before"`
	Now    []Dimension `json:"now"`
	After  []Dimension `json:"after"`
}

// 获取性格纬度列表
func (c *Client) ListPersonalityDimensions(ctx context.Context, params *ListPersonalityDimensionsParams) (*ListPersonalityDimensionsResult, error) {
	bytes, err := json.Marshal(params)
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
		SetPath("/personality/dimensions").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(ListPersonalityDimensionsResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
