package recombee

import (
	"fmt"
	"net/http"
	"strings"
	"testing"
	//"github.com/mavvverick/recombee"
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

type itemData struct {
	User          string   `json:"user"`
	Script        string   `json:"script,omitempty"`
	Cat           []string `json:"cat,omitempty"`
	Desc          string   `json:"desc,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Img           string   `json:"img,omitempty"`
	CascadeCreate bool     `json:"!cascadeCreate,omitempty"`
}

func TestAction_SETProp(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	/*
		itemId, userId, tags, desc, script, topic, cat, imgs
	*/

	itemData := itemData{
		User:          "123",
		Desc:          "yahoo!",
		Tags:          strings.Split("alpha,beta", ","),
		Img:           fmt.Sprintf("https://playy-test.s3.ap-south-1.amazonaws.com/%v/intro.jpg", "12"),
		CascadeCreate: true,
	}

	item := &Item{
		ID: "12",
	}

	resp, _, err := client.Item.SetProp(ctx, item, itemData)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	fmt.Println(resp)
}
