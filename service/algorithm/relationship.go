package algorithm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/muwanyou/nonba-sdk/enum"
)

// 关系参数
type ListRelationshipsInput struct {
	SubjectFamilyName string    `json:"subject_family_name"`
	SubjectGivenName  string    `json:"subject_given_name"`
	SubjectSex        enum.Sex  `json:"subject_sex"`
	SubjectBirthday   time.Time `json:"subject_birthday"`
	ObjectFamilyName  string    `json:"object_family_name"`
	ObjectGivenName   string    `json:"object_given_name"`
	ObjectSex         enum.Sex  `json:"object_sex"`
	ObjectBirthday    time.Time `json:"object_birthday"`
}

// 关系结果
type ListRelationshipsOutput struct {
	Items []string `json:"items"`
}

// 关系列表
func (c *Client) ListRelationships(ctx context.Context, input *ListRelationshipsInput) (*ListRelationshipsOutput, error) {
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
		SetPath("/relationships").
		SetQuery(query)
	response := NewResponse()
	err = c.Send(request, response)
	if err != nil {
		return nil, err
	}
	output := new(ListRelationshipsOutput)
	err = json.Unmarshal(response.GetBody(), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
