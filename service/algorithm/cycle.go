package algorithm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

// 周期参数
type GetCycleParam struct {
	TimeUnit   enum.TimeUnit      `json:"time_unit"`
	Polarity   enum.CyclePolarity `json:"polarity"`
	FamilyName string             `json:"family_name"`
	GivenName  string             `json:"given_name"`
	Sex        enum.Sex           `json:"sex"`
	Birthday   time.Time          `json:"birthday"`
	Datetime   time.Time          `json:"datetime"`
}

// 周期结果
type GetCycleResult struct {
	ID     int64  `json:"id,string"`
	Name   string `json:"name"`
	Impact string `json:"impact"`
}

// 获取周期
func (c *Client) GetCycle(ctx context.Context, param *GetCycleParam) (*GetCycleResult, error) {
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
		SetPath(fmt.Sprintf("/cycles/%s/%s", param.TimeUnit, param.Polarity)).
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(GetCycleResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 周期纬度列表参数
type ListCycleDimensionsParam struct {
	TimeUnit   enum.TimeUnit      `json:"time_unit"`
	Polarity   enum.CyclePolarity `json:"polarity"`
	FamilyName string             `json:"family_name"`
	GivenName  string             `json:"given_name"`
	Sex        enum.Sex           `json:"sex"`
	Birthday   time.Time          `json:"birthday"`
	Datetime   time.Time          `json:"datetime"`
	Height     int16              `json:"height,string"`
	Weight     int16              `json:"weight,string"`
}

// 周期纬度列表结果
type ListCycleDimensionsResult struct {
	Items []*Dimension `json:"items"`
}

// 获取周期纬度列表
func (c *Client) ListCycleDimensions(ctx context.Context, param *ListCycleDimensionsParam) (*ListCycleDimensionsResult, error) {
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
		SetPath(fmt.Sprintf("/cycles/%s/%s/dimensions", param.TimeUnit, param.Polarity)).
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(ListCycleDimensionsResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
