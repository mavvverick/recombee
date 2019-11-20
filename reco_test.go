package recombee

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAction_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	u := User{ID: "1"}
	l := logics["recombee:default"]

	recoms, _, err := client.Reco.GetPreset(ctx, u, l)
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

	u := User{ID: "1"}
	opts := &ListOptions{
		Count:        10,
		RotationRate: 1,
	}
	recoms, _, err := client.Reco.ItemsToUser(ctx, u, opts)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	ids := []string{}

	for _, v := range recoms.Recomms {
		ids = append(ids, v.ID)
	}
	fmt.Println(ids)
	fmt.Println(recoms)
}
