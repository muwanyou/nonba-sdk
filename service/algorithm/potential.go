package algorithm

import (
	"context"
	"encoding/json"

	"github.com/muwanyou/nonba-sdk/enum"
)

type PotentialElement struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name"`
	Letter      string `json:"letter"`
	Description string `json:"description"`
	Positive    string `json:"positive"`
	Negative    string `json:"negative"`
}

// 潜能参数
type GetPotentialParam struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
}

// 潜能结果
type GetPotentialResult struct {
	ID       int64               `json:"id,string"`
	Elements []*PotentialElement `json:"elements"`
	Keywords []string            `json:"keywords"`
}

// 获取潜能
func (c *Client) GetPotential(ctx context.Context, param *GetPotentialParam) (*GetPotentialResult, error) {
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
		SetPath("/potential").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(GetPotentialResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 潜能纬度列表参数
type ListPotentialDimensionsParam struct {
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
}

// 潜能纬度列表结果
type ListPotentialDimensionsResult struct {
	Items []*Dimension `json:"items"`
}

// 获取潜能纬度列表
func (c *Client) ListPotentialDimensions(ctx context.Context, param *ListPotentialDimensionsParam) (*ListPotentialDimensionsResult, error) {
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
		SetPath("/potential/dimensions").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(ListPotentialDimensionsResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
