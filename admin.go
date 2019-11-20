package recombee

import (
	"context"
	"fmt"
	"net/http"
)

//BatchService handle batch comm to
// recombee API /{databaseId}/batch/
type AdminService interface {
	Delete(context.Context) (*Response, error)
}

//BatchServiceOp handle batch comm to
// recombee API /{databaseId}/batch/
type AdminServiceOp struct {
	client *Client
}

var _ AdminService = &AdminServiceOp{}

//Request structure for batch

//Delete database request to recombee
func (s *AdminServiceOp) Delete(ctx context.Context) (*Response, error) {
	path := fmt.Sprintf("/%v/?", db)
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
