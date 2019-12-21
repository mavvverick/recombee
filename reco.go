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
	GetPreset(context.Context, *User, *ListOptions) (*RecoRoot, *Response, error)
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

type Settings struct {
	MaxAge int64 `json:"maxAge"`
}

type Logic struct {
	Name     string    `json:"name"`
	Settings *Settings `json:"settings"`
}

// ListOptions for recommendations
type ListOptions struct {
	Count              int64         `json:"count"`
	Filter             string        `json:"filter,omitempty"`
	Booster            string        `json:"booster,omitempty"`
	CascadeCreate      bool          `json:"cascadeCreate,omitempty"`
	Scenario           string        `json:"scenario,omitempty"`
	Logic              *Logic        `json:"logic,omitempty"`
	ReturnProperties   bool          `json:"returnProperties,omitempty"`
	IncludedProperties []interface{} `json:"includedProperties,omitempty"`
	Diversity          float32       `json:"diversity,omitempty"`
	MinRelevance       string        `json:"minRelevance,omitempty"`
	RotationRate       float32       `json:"rotationRate"`
	RotationTime       float64       `json:"rotationTime"`
}

var _ RecoService = &RecoServiceOp{}

// GetPreset recommendation by keywords.
func (s *RecoServiceOp) GetPreset(ctx context.Context, u *User, opt *ListOptions) (*RecoRoot, *Response, error) {
	iskey := validate(opt.Logic.Name)
	if iskey != true {
		return nil, nil, errors.New("Provide valid recombee logic")
	}

	if opt.Count == 0 {
		opt.Count = 10
	}

	path := fmt.Sprintf("/%v/recomms/users/%v/items/?", db, u.ID)
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
	if opt.Count == 0 {
		opt.Count = 10
	}
	return reco(ctx, s, path, opt)
}

// UsersToUser recommendations
func (s *RecoServiceOp) UsersToUser(ctx context.Context, u *User, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/users/%v/users/?", db, u.ID)
	if opt.Count == 0 {
		opt.Count = 10
	}
	return reco(ctx, s, path, opt)
}

// ItemsToItem recommendations
func (s *RecoServiceOp) ItemsToItem(ctx context.Context, i *Item, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/items/%v/items/?", db, i.ID)
	if opt.Count == 0 {
		opt.Count = 10
	}
	return reco(ctx, s, path, opt)
}

// UsersToItem recommendations
func (s *RecoServiceOp) UsersToItem(ctx context.Context, i *Item, opt *ListOptions) (*RecoRoot, *Response, error) {
	path := fmt.Sprintf("/%v/recomms/items/%v/users/?", db, i.ID)
	if opt.Count == 0 {
		opt.Count = 10
	}
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
