package recombee

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestAction_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	u := &User{ID: "1"}
	opts := &ListOptions{
		Count:        10,
		RotationRate: 1,
		Logic: &Logic{
			Name: logics["recombee:default"],
		},
	}

	recoms, _, err := client.Reco.GetPreset(ctx, u, opts)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	fmt.Println(recoms)
}

func TestAction_ItemToUser(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	u := &User{ID: "1"}
	opts := &ListOptions{
		Count:        2,
		RotationRate: 1,
		RotationTime: 3.154e+7,
	}

	//opts.Filter = `"exp3min" in 'cat'`,
	tMinus2Days := time.Now().Add(-24 * 2 * time.Hour).Unix()
	tMinus5Days := time.Now().Add(-24 * 5 * time.Hour).Unix()
	tMinus7Days := time.Now().Add(-24 * 7 * time.Hour).Unix()
	fmt.Println(tMinus2Days)
	opts.Booster = fmt.Sprintf(`if 'when' >= %v then 3 else if ('when' <= %v and 'when' >= %v) then 2.5 else if ('when' <= %v and 'when' >= %v) then 1.5 else 1`, tMinus2Days, tMinus2Days, tMinus5Days, tMinus5Days, tMinus7Days)
	//opts.Booster = fmt.Sprintf(`if ('when' >= %v and "exp3min" not in 'cat') then 3 else 1.5`, tMinus2Days)
	recoms, resp, err := client.Reco.ItemsToUser(ctx, u, opts)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	fmt.Println(resp)

	ids := []string{}

	for _, v := range recoms.Recomms {
		ids = append(ids, v.ID)
	}
	fmt.Println(ids)
	fmt.Println(recoms)
}
