package transport

import "github.com/gophers0/users/internal/model"

type SearchRequest struct {
	Login  string `json:"login"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

func (req *SearchRequest) Validate() error {
	return nil
}

type SearchResponse struct {
	BaseResponse
	Count   int           `json:"count"`
	Records []*model.User `json:"records"`
}
