package recombee

import (
	"context"
	"net/http"
)

//BatchService handle batch comm to
// recombee API /{databaseId}/batch/
type BatchService interface {
	Post(context.Context, Batches) (*Response, error)
}

//BatchServiceOp handle batch comm to
// recombee API /{databaseId}/batch/
type BatchServiceOp struct {
	client *Client
}

var _ BatchService = &BatchServiceOp{}

type Batches struct {
	Requests []Request `json:"requests"`
}

//Request structure for batch
type Request struct {
	Method string      `json:"method"`
	Path   string      `json:"path"`
	Params interface{} `json:"params"`
}

//Post batch request to recombee
func (s *BatchServiceOp) Post(ctx context.Context, b Batches) (*Response, error) {
	path := "/totality-dev/batch/"
	url := GenURL(path)
	req, err := s.client.NewRequest(ctx, http.MethodPut, url, b.Requests)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}
