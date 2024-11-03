package entity

import "errors"

type APIInfo struct {
	Method string
	Path   string
	Query  map[string]string
	Body   []byte
}

func NewAPIInfo(
	method,
	path string,
	query map[string]string,
	body []byte,
) (*APIInfo, error) {
	if method != "GET" && method != "POST" {
		return nil, errors.New("invalid method")
	}
	if method == "POST" && body == nil {
		return nil, errors.New("body is required")
	}

	return &APIInfo{
		Method: method,
		Path:   path,
		Query:  query,
		Body:   body,
	}, nil
}
