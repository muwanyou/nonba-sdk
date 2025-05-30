package algorithm_test

import (
	"context"
	"testing"
	"time"

	"github.com/muwanyou/nonba-sdk/core"
	"github.com/muwanyou/nonba-sdk/enum"
	"github.com/muwanyou/nonba-sdk/service/algorithm"
)

func TestListRelationships(t *testing.T) {
	ctx := context.Background()
	credential := core.NewCredential("3b6d4e442cf7422cb8431419068802f6", "1", "PJ8^yl*s5jG*Yvhlgm5N!u0Suljnx^K&")
	client := algorithm.NewClient(credential)
	Output, err := client.ListRelationships(ctx, &algorithm.ListRelationshipsInput{
		SubjectFamilyName: "颜",
		SubjectGivenName:  "吉灿",
		SubjectSex:        enum.SexMale,
		SubjectBirthday:   time.Now(),
		ObjectFamilyName:  "张",
		ObjectGivenName:   "三丰",
		ObjectSex:         enum.SexMale,
		ObjectBirthday:    time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(Output)
}
