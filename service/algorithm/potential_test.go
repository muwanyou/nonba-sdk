package algorithm_test

import (
	"context"
	"testing"

	"github.com/muwanyou/nonba-sdk/core"
	"github.com/muwanyou/nonba-sdk/service/algorithm"
)

func TestGetPotential(t *testing.T) {
	ctx := context.Background()
	credential := core.NewCredential("3b6d4e442cf7422cb8431419068802f6", "1", "PJ8^yl*s5jG*Yvhlgm5N!u0Suljnx^K&")
	client := algorithm.NewClient(credential)
	result, err := client.GetPotential(ctx, &algorithm.GetPotentialParam{
		FamilyName: "张",
		GivenName:  "三丰",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestListPotentialDimensions(t *testing.T) {
	ctx := context.Background()
	credential := core.NewCredential("3b6d4e442cf7422cb8431419068802f6", "1", "PJ8^yl*s5jG*Yvhlgm5N!u0Suljnx^K&")
	client := algorithm.NewClient(credential)
	result, err := client.ListPotentialDimensions(ctx, &algorithm.ListPotentialDimensionsParam{
		FamilyName: "颜",
		GivenName:  "吉灿",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
