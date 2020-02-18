package recombee

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var boostScheme = fmt.Sprintf(`if 'when' >= %v then 5 else if ('when' < %v and 'when' >= %v) then 4 else if ('when' < %v and 'when' >= %v) then 3.5 else if ('when' < %v and 'when' >= %v) then 3 else if ('when' < %v and 'when' >= %v) then 2 else 1 `, day(2), day(2), day(5), day(5), day(10), day(10), day(30), day(30), day(60))

func TestAction_Get(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	u := &User{ID: "2"}
	opts := &ListOptions{
		Count: 3,
		//Booster: fmt.Sprintf("if 'when' >= %v then 5 else if ('when' < %v and 'when' >=  %v) then 2 else 0", day(2), day(4), day(6)),
		//Booster: fmt.Sprintf("if 'when' >= %v then 100000 else 1", day(10)),
		//Filter:  " 'when' >= now() - (25 * 24 * 60 * 60)",
		//Booster: "if 'when' >= now() - (5 * 24 * 60 * 60) then 10.00 else 1",
		//Booster:            fmt.Sprintf(`if 'tags' in {"expectations"} then 10000000 else 1`),
		RotationRate:     1,
		ReturnProperties: true,
		//IncludedProperties: "when,user",
		IncludedProperties: "tags,username,rating,tracks,user",
		RotationTime:       30000,
		CascadeCreate:      true,
	}

	//opts.Booster = getRecencyBooster()
	//opts.Diversity = 1.0
	// opts.Logic = &Logic{
	// 	Name: "recombee:popular",
	// }

	fmt.Println(opts)
	//recoms, _, err := client.Reco.GetPreset(ctx, u, opts)
	recoms, _, err := client.Reco.ItemsToUser(ctx, u, opts)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	for _, data := range recoms.Recomms {
		Data := data.Values.(map[string]interface{})
		//fmt.Println(data.ID, Data["user"], time.Unix(int64(Data["when"].(float64)), 0))
		fmt.Println(data.ID, Data["user"])
	}
}

func TestAction_GetAsync(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	u := &User{ID: "2"}
	opts := &ListOptions{
		Count:              1,
		RotationRate:       1,
		ReturnProperties:   true,
		IncludedProperties: "tags, username,rating,tracks,user",
		RotationTime:       30000,
		CascadeCreate:      true,
	}

	c := make(chan *RecoRoot, 2)

	go client.Reco.ItemsToUserAsync(ctx, u, opts, c)
	go client.Reco.ItemsToUserAsync(ctx, u, opts, c)

	a := []*RecoRoot{<-c, <-c}
	for _, list := range a {
		fmt.Println(list)
		for _, data := range list.Recomms {
			Data := data.Values.(map[string]interface{})
			fmt.Println(data.ID, Data["user"])
		}
	}

	defer close(c)
}

func getRecencyBooster() string {
	tMinus2Days := time.Now().Add(-24 * 2 * time.Hour).Unix()
	tMinus5Days := time.Now().Add(-24 * 5 * time.Hour).Unix()
	tMinus7Days := time.Now().Add(-24 * 7 * time.Hour).Unix()
	return fmt.Sprintf(`if 'when' >= %v then 3 else if ('when' <= %v and 'when' >= %v) then 2.5 else if ('when' <= %v and 'when' >= %v) then 1.5 else 1`, tMinus2Days, tMinus2Days, tMinus5Days, tMinus5Days, tMinus7Days)
}

func day(d int64) string {
	return fmt.Sprintf(`now() - %v`, (d * 86400))
}

func TestAction_ItemToUser(t *testing.T) {
	setup()
	defer teardown()
	mux.HandleFunc("/v1/reco/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodPost)
	})

	//Get user view count from redis:: move to client, ask them to fetch at the begining of session
	//if < 0: recommend popular
	//if > 10: recomment default
	//if seen too many from same superscript filter topic
	//balance algorithm :: when to add burst of random and popular video again in feed
	//get relation between script / similarities

	u := &User{ID: "1"}
	opts := &ListOptions{
		Count:        2,
		RotationRate: 1,
		//RotationTime: 3.154e+7,
	}

	opts.Diversity = 0.4

	//TODO check behaviour with interaction
	// opts.MinRelevance = "medium"

	// //opts.Filter = `"exp3min" in 'cat'`,
	// tMinus2Days := time.Now().Add(-24 * 2 * time.Hour).Unix()
	// tMinus5Days := time.Now().Add(-24 * 5 * time.Hour).Unix()
	// tMinus7Days := time.Now().Add(-24 * 7 * time.Hour).Unix()
	// fmt.Println(tMinus2Days)
	// opts.Booster = fmt.Sprintf(`if 'when' >= %v then 3 else if ('when' <= %v and 'when' >= %v) then 2.5 else if ('when' <= %v and 'when' >= %v) then 1.5 else 1`, tMinus2Days, tMinus2Days, tMinus5Days, tMinus5Days, tMinus7Days)

	opts.Booster = boostScheme
	fmt.Println(opts)
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
