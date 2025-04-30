package algorithm

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

type GetCharacterParams struct {
	FamilyName string    `json:"family_name"`
	GivenName  string    `json:"given_name"`
	Sex        enum.Sex  `json:"sex"`
	Birthday   time.Time `json:"birthday"`
	Year       int16     `json:"year,string"`
	Height     int16     `json:"height,string"`
	Weight     int16     `json:"weight,string"`
}

type GetCharacterResult struct {
	Before map[string][]string `json:"before"`
	Now    map[string][]string `json:"now"`
	After  map[string][]string `json:"after"`
}

func (c *Client) GetCharacter(ctx context.Context, params *GetCharacterParams) (*GetCharacterResult, error) {
	bytes, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	query := make(map[string]string)
	err = json.Unmarshal(bytes, &query)
	if err != nil {
		return nil, err
	}
	fmt.Println(query)
	request := NewRequest()
	request.SetContext(ctx).
		SetMethod(enum.MethodGet).
		SetPath("/character").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	result := new(GetCharacterResult)
	err = json.Unmarshal(response.GetBody(), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
