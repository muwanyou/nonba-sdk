package algorithm_test

import (
	"context"
	"testing"
	"time"

	"github.com/muwanyou/nonba-sdk/core"
	"github.com/muwanyou/nonba-sdk/enum"
	"github.com/muwanyou/nonba-sdk/service/algorithm"
)

func TestGetCycle(t *testing.T) {
	ctx := context.Background()
	credential := core.NewCredential("3b6d4e442cf7422cb8431419068802f6", "1", "PJ8^yl*s5jG*Yvhlgm5N!u0Suljnx^K&")
	client := algorithm.NewClient(credential)
	result, err := client.GetCycle(ctx, &algorithm.GetCycleParam{
		TimeUnit:   enum.TimeUnitYear,
		Polarity:   enum.CyclePolarityPositive,
		FamilyName: "颜",
		GivenName:  "吉灿",
		Sex:        enum.SexMale,
		Birthday:   time.Now(),
		Datetime:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}

func TestListCycleDimensions(t *testing.T) {
	ctx := context.Background()
	credential := core.NewCredential("3b6d4e442cf7422cb8431419068802f6", "1", "PJ8^yl*s5jG*Yvhlgm5N!u0Suljnx^K&")
	client := algorithm.NewClient(credential)
	result, err := client.ListCycleDimensions(ctx, &algorithm.ListCycleDimensionsParam{
		TimeUnit:   enum.TimeUnitYear,
		Polarity:   enum.CyclePolarityPositive,
		FamilyName: "颜",
		GivenName:  "吉灿",
		Sex:        enum.SexMale,
		Birthday:   time.Now(),
		Datetime:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(result)
}
