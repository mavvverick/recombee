package recombee

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

// RecoService handles communction with recommendation related methods of the
// Recombee API:/{databaseId}/recomms/users/{userId}/items/
type RecoService interface {
	GetPreset(context.Context, *User, string, *ListOptions) (*RecoRoot, *Response, error)
	ItemsToUser(context.Context, *User, *ListOptions) (*RecoRoot, *Response, error)
	UsersToUser(context.Context, *User, *ListOptions) (*RecoRoot, *Response, error)
	ItemsToItem(context.Context, *Item, *ListOptions) (*RecoRoot, *Response, error)
	UsersToItem(context.Context, *Item, *ListOptions) (*RecoRoot, *Response, error)
}

// RecoServiceOp handles communction with recommendation related methods of the
// Recombee API:/{databaseId}/recomms/users/{userId}/items/
type RecoServiceOp struct {
	client *Client
}

type Recomm struct {
	ID string `json:"id"`
}

type RecoRoot struct {
	RecommID string   `json:"recommId"`
	Recomms  []Recomm `json:"recomms"`
}

// ListOptions for recommendations
type ListOptions struct {
	Count              int32         `json:"count"`
	Filter             string        `json:"filter,omitempty"`
	Booster            string        `json:"booster,omitempty"`
	CascadeCreate      bool          `json:"cascadeCreate,omitempty"`
	Scenario           string        `json:"scenario,omitempty"`
	Logic              interface{}   `json:"logic,omitempty"`
	ReturnProperties   bool          `json:"returnProperties,omitempty"`
	IncludedProperties []interface{} `json:"includedProperties,omitempty"`
	Diversity          int64         `json:"diversity,omitempty"`
	MinRelevance       string        `json:"minRelevance,omitempty"`
	RotationRate       float32       `json:"rotationRate,omitempty"`
	RotationTime       float32       `json:"rotationTime,omitempty"`
}

var _ RecoService = &RecoServiceOp{}

// GetPreset recommendation by keywords.
func (s *RecoServiceOp) GetPreset(ctx context.Context, u *User, l string, opt *ListOptions) (*RecoRoot, *Response, error) {
	iskey := validate(l)
	if iskey != true {
		return nil, nil, errors.New("Provide valid recombee logic")
	}
	count := 10
	path := fmt.Sprintf("/%v/recomms/users/%v/items/?count=%v&logic=%v&", db, u.ID, count, l)
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodPost, url, opt)
	if err != nil {
		return nil, nil, err
	}

	root := new(RecoRoot)

	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// ItemsToUser recommendations
func (s *RecoServiceOp) ItemsToUser(ctx context.Context, u *User, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/users/%v/items/?", db, u.ID)
	return reco(ctx, s, path, opt)
}

// UsersToUser recommendations
func (s *RecoServiceOp) UsersToUser(ctx context.Context, u *User, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/users/%v/users/?", db, u.ID)
	return reco(ctx, s, path, opt)
}

// ItemsToItem recommendations
func (s *RecoServiceOp) ItemsToItem(ctx context.Context, i *Item, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/items/%v/items/?", db, i.ID)
	return reco(ctx, s, path, opt)
}

// UsersToItem recommendations
func (s *RecoServiceOp) UsersToItem(ctx context.Context, i *Item, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/items/%v/users/?", db, i.ID)
	return reco(ctx, s, path, opt)
}

func reco(ctx context.Context, s *RecoServiceOp, path string, opt *ListOptions) (*RecoRoot, *Response, error) {
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodPost, url, opt)
	if err != nil {
		return nil, nil, err
	}

	root := new(RecoRoot)
	resp, err := s.client.Do(ctx, req, root)

	if err != nil {
		return nil, resp, err
	}

	return root, resp, nil
}

// func (s *RecoServiceOp) ItemToUser(ctx context.Context, u User, opt *ListOptions) (*RecoRoot, *Response, error) {
// 	path := fmt.Sprintf("/%v/recomms/users/%v/items/?", db, u.ID)
// 	opt.Count = 10
// 	url := GenURL(path)

// 	req, err := s.client.NewRequest(ctx, http.MethodPost, url, opt)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	root := new(RecoRoot)

// 	resp, err := s.client.Do(ctx, req, root)
// 	if err != nil {
// 		return nil, resp, err
// 	}

// 	return root, resp, nil
// }

func validate(s string) bool {
	_, ok := logics[s]
	return ok
}
