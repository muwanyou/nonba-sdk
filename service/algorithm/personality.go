package algorithm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

type Personality struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Positive    string `json:"positive"`
	Negative    string `json:"negative"`
	Solve       string `json:"solve"`
	Description string `json:"description"`
}

// 性格列表参数
type ListPersonalitiesParam struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
}

// 性格列表结果
type ListPersonalitiesResult struct {
	Items []*Personality `json:"items"`
}

// 获取性格列表
func (c *Client) ListPersonalities(ctx context.Context, params *ListPersonalitiesParam) (*ListPersonalitiesResult, error) {
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
type ListPersonalityDimensionsParam struct {
	Timeline   enum.Timeline `json:"timeline"`
	FamilyName string        `json:"family_name"`
	GivenName  string        `json:"given_name"`
	Sex        enum.Sex      `json:"sex"`
	Birthday   time.Time     `json:"birthday"`
	Datetime   time.Time     `json:"datetime"`
	Height     int16         `json:"height"`
	Weight     int16         `json:"weight"`
}

// 性格纬度列表结果
type ListPersonalityDimensionsResult struct {
	Items []Dimension `json:"items"`
}

// 获取性格纬度列表
func (c *Client) ListPersonalityDimensions(ctx context.Context, param *ListPersonalityDimensionsParam) (*ListPersonalityDimensionsResult, error) {
	bytes, err := json.Marshal(param)
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
		SetPath(fmt.Sprintf("/personality/dimensions/%s", param.Timeline)).
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
