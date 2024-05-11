package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func makeGetPersonEndpoint(s Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getRequest)
		person, err := s.GetPersonInfo(req.PersonId)
		if err != nil {
			return nil, err
		}
		return person, nil
	}
}
