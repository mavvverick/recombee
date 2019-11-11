package recombee

import (
	"context"
	"fmt"
	"net/http"
)

// ItemService handles communction with action related methods
type ItemService interface {
	Post(context.Context, Item) (*Response, error)
	List(context.Context) (*ItemsRoot, *Response, error)
	Delete(context.Context, Item) (*Response, error)
	AddProp(context.Context, ItemProperty) (*Response, error)
	DeleteProp(context.Context, ItemProperty) (*Response, error)
	ListProp(context.Context) (*Response, error)

	SetProp(context.Context, Item, interface{}) (*Response, error)
	GetProp(context.Context, Item) (*Response, error)
}

// ItemServiceOp handles communition with the items action related methods
type ItemServiceOp struct {
	client *Client
}

var _ ItemService = &ItemServiceOp{}

type Item struct {
	ID string
}

type ItemProperty struct {
	name string
	typ  string
}

// ItemsRoot represents a Recombee Item
type ItemsRoot []string

//Post items to recombee
func (s *ItemServiceOp) Post(ctx context.Context, i Item) (*Response, error) {
	path := "/totality-dev/items/" + i.ID
	url := GenURL(path)
	req, err := s.client.NewRequest(ctx, http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//Delete items to recombee
func (s *ItemServiceOp) Delete(ctx context.Context, i Item) (*Response, error) {
	path := "/totality-dev/items/" + i.ID
	url := GenURL(path)
	req, err := s.client.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//List items to recombee
func (s *ItemServiceOp) List(ctx context.Context) (*ItemsRoot, *Response, error) {
	//TODO filter options in url
	path := "/totality-dev/items/list/"
	url := GenURL(path)
	req, err := s.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	root := new(ItemsRoot)

	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	return root, resp, err
}

//AddProp to recombee items
func (s *ItemServiceOp) AddProp(ctx context.Context, i ItemProperty) (*Response, error) {
	path := fmt.Sprintf("/totality-dev/items/properties/%v?type=%v", i.name, i.typ)
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodPut, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//DeleteProp of recombee items
func (s *ItemServiceOp) DeleteProp(ctx context.Context, i ItemProperty) (*Response, error) {
	path := fmt.Sprintf("/totality-dev/items/properties/%v", i.name)
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//ListProp of recombee items
func (s *ItemServiceOp) ListProp(ctx context.Context) (*Response, error) {
	path := "/totality-dev/items/properties/list/"
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//SetProp of recombee items
func (s *ItemServiceOp) SetProp(ctx context.Context, i Item, m interface{}) (*Response, error) {
	path := "/totality-dev/items/" + i.ID
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodPost, url, m)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

//GetProp of recombee items
func (s *ItemServiceOp) GetProp(ctx context.Context, i Item) (*Response, error) {
	path := "/totality-dev/items/" + i.ID
	url := GenURL(path)

	req, err := s.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
