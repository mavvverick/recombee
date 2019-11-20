package recombee

import (
	"context"
	"fmt"
	"net/http"
)

//BatchService handle batch comm to
// recombee API /{databaseId}/batch/
type BatchService interface {
	Post(context.Context, *Batches) (*BatchesRoot, *Response, error)
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
//NOTE remeber to use CascadeCreate in Params  `json:"!cascadeCreate,omitempty"`
type Request struct {
	Method string      `json:"method"`
	Path   string      `json:"path"`
	Params interface{} `json:"params"`
}

type batchRoot struct {
	Code int64       `json:"code"`
	JSON interface{} `json:"json"`
}

type BatchesRoot []batchRoot

//Post batch request to recombee
func (s *BatchServiceOp) Post(ctx context.Context, bat *Batches) (*BatchesRoot, *Response, error) {
	path := fmt.Sprintf("/%v/batch/?", db)
	url := GenURL(path)
	req, err := s.client.NewRequest(ctx, http.MethodPost, url, bat)
	if err != nil {
		return nil, nil, err
	}

	root := new(BatchesRoot)

	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
