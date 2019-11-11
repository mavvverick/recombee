package recombee

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

// func TestAction_List(t *testing.T) {
// 	setup()
// 	defer teardown()
// 	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
// 		testMethod(t, r, http.MethodPost)
// 	})

// 	items, _, err := client.Item.List(ctx)

// 	if err != nil {
// 		t.Fatalf("unexpected error: %s", err)
// 	}

// 	fmt.Println(*items)
// }
