package recombee

import (
	"context"
	"fmt"
	"net/http"
)

// UserService handles communction with action related methods
type UserService interface {
	Post(context.Context, User) (*Response, error)
}

type User struct {
	ID string
}

// UserServiceOp handles communition with the items action related methods
type UserServiceOp struct {
	client *Client
}

var _ UserService = &UserServiceOp{}

// User represents a Recombee User
type UsersRoot []string

//Post items to recombee
func (s *UserServiceOp) Post(ctx context.Context, u User) (*Response, error) {
	path := fmt.Sprintf("/%v/users/%v?", db, u.ID)

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

func (s *UserServiceOp) Delete(ctx context.Context, u User) (*Response, error) {
	path := fmt.Sprintf("/%v/users/%v", db, u.ID)
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

func (s *UserServiceOp) List(ctx context.Context) (*UsersRoot, *Response, error) {
	//TODO filter options in url
	path := fmt.Sprintf("/%v/users/list/", db)
	url := GenURL(path)
	req, err := s.client.NewRequest(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	root := new(UsersRoot)

	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	return root, resp, err
}

func (s *UserServiceOp) Set(ctx context.Context, u User, m interface{}) (*Response, error) {
	path := fmt.Sprintf("/%v/users/%v", db, u.ID)
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
