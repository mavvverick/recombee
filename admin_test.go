package recombee

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAction_ResetDB(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	resp, err := client.Admin.Delete(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	fmt.Println(resp)
}
