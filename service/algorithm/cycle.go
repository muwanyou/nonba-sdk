package algorithm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

// 周期参数
type GetCycleInput struct {
	TimeUnit   enum.TimeUnit      `json:"time_unit"`
	Polarity   enum.CyclePolarity `json:"polarity"`
	FamilyName string             `json:"family_name"`
	GivenName  string             `json:"given_name"`
	Sex        enum.Sex           `json:"sex"`
	Birthday   time.Time          `json:"birthday"`
	Datetime   time.Time          `json:"datetime"`
}

// 周期结果
type GetCycleOutput struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Impact      string `json:"impact"`
}

// 获取周期
func (c *Client) GetCycle(ctx context.Context, input *GetCycleInput) (*GetCycleOutput, error) {
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
		SetPath(fmt.Sprintf("/cycles/%s/%s", input.TimeUnit, input.Polarity)).
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	output := new(GetCycleOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// 周期纬度列表参数
type ListCycleDimensionsInput struct {
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
type ListCycleDimensionsOutput struct {
	Items []*Dimension `json:"items"`
}

// 获取周期纬度列表
func (c *Client) ListCycleDimensions(ctx context.Context, input *ListCycleDimensionsInput) (*ListCycleDimensionsOutput, error) {
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
		SetPath(fmt.Sprintf("/cycles/%s/%s/dimensions", input.TimeUnit, input.Polarity)).
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	output := new(ListCycleDimensionsOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
