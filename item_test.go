package recombee

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"testing"
	"time"
)

type itemData struct {
	User          string   `json:"user"`
	Script        string   `json:"script,omitempty"`
	Cat           []string `json:"cat,omitempty"`
	Desc          string   `json:"desc,omitempty"`
	When          int64    `json:"when,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Img           string   `json:"img,omitempty"`
	CascadeCreate bool     `json:"!cascadeCreate,omitempty"`
}

func TestAction_Post(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	i := &Item{ID: "computer-0"}

	resp, err := client.Item.Post(ctx, i)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	fmt.Println(resp)
}

func TestAction_ListAndUpdateItems(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	itemIDS, _, err := client.Item.List(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	rand.Seed(time.Now().UnixNano())
	tMinus60 := time.Now().Add(-24 * 21 * time.Hour)

	for _, itemID := range *itemIDS {
		count := rand.Intn(21)
		time := tMinus60.Add(24 * time.Duration(count) * time.Hour)
		fmt.Println(count, time)
		fmt.Println("Updating... ", itemID)

		item := &Item{
			ID: itemID,
		}

		itemData := &itemData{
			When: time.Unix(),
		}

		ok, _, err := client.Item.SetProp(ctx, item, *itemData)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		fmt.Println(*ok)
	}

}

func TestAction_Delete(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	s := "aMGYUzAwO0XWByf,99a9w6-Wg,KNIN8eaZg,nSaBueaZg,NVU0s6-ZR,of7RU6-WR,QbqiggBZg,qtZdggBWg,_Bbz_6-Wg,IAH8kbH2Z0NfiH7,cQzPa6-Zg,EoYB93-ZR,EwoWa6-ZR,fknjo6-WR,gr9C-eaZR,hF1I9qaWR,LnC893aWg,t33Or3aZR,VotqXq-Wg,TLrHOMRdzfEcqPm,1_3nDeaZg,2tZeW6aWR,4EVfU6aZR,5HS3ue-Zg,7FxaNe-Wg,7VCdo6-Wg,8Pe-l6aZg,8qC4s6-Wg,92zuQ6-WR,9MYFY6-ZR,A4nwx6aWg,D2RAYe-ZR,dEy9re-ZR,flLddiBZg,I0_ZE6aWg,iAhvPe-ZR,iljrCeaZR,Jgzvf6aZg,JtqpNmBZg,MHWMf6-ZR,nov68eaWg,oYTyeeaZR,plGoXeaZg,Q7cBJ6aZR,QaZgdmBWR,QKAQ16aWg,QW6JI6aZg,RSNV26-Wg,RvPIpe-WR,sUm4i6aZR,tULst6aWg,vwkJx6-Wg,wA7SYeaZg,WWRRfeaWR,XUdpS6aWg,xwBpvmfWR,yc2BJe-Wg,YFaVE6-WR,YVHQSeaWR,Z8n1diBWg,Z9jL_eaWg,ZIYpJe-WR,FwEWy3aWg"

	ids := strings.Split(s, ",")

	for _, id := range ids {
		i := &Item{ID: id}
		resp, err := client.Item.Delete(ctx, i)
		if err != nil {
			fmt.Println(err)
			// t.Fatalf("unexpected error: %s", err)
		}

		fmt.Println(resp)
	}
}

func TestAction_AddProp(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/items/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	properties := []ItemProperty{
		// ItemProperty{Name: "tags", Type: "set"},
		// ItemProperty{Name: "cat", Type: "set"},
		// ItemProperty{Name: "desc", Type: "string"},
		// ItemProperty{Name: "msg", Type: "string"},
		// ItemProperty{Name: "user", Type: "string"},
		// ItemProperty{Name: "img", Type: "image"},
		ItemProperty{Name: "when", Type: "timestamp"},
	}

	for _, prop := range properties {
		items, err := client.Item.AddProp(ctx, &prop)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		fmt.Println(*items)
	}

	//items, err := client.Item.DeleteProp(ctx, prop)
	//items, _, err := client.Item.ListProp(ctx)

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

	fmt.Println(*resp)
}
