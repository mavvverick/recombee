package recombee

import (
	"fmt"
	"net/http"
	"testing"
)

// import (
// 	"fmt"
// 	"net/http"
// 	"testing"
// )

// func TestAction_Post(t *testing.T) {
// 	setup()
// 	defer teardown()
// 	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, http.MethodPost)
// 	})

// 	i := Item{ID: "computer-0"}

// 	resp, err := client.Item.Post(ctx, i)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %s", err)
// 	}

// 	fmt.Println(resp)
// }

func TestAction_AddProp(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	/*
		itemId, userId, tags, desc, script, topic, cat, imgs
	*/

	// prop := ItemProperty{
	// 	Name: "img",
	// 	Type: "image",
	// }
	// items, err := client.Item.AddProp(ctx, prop)
	//items, err := client.Item.DeleteProp(ctx, prop)

	items, _, err := client.Item.ListProp(ctx)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	fmt.Println(*items)
}
