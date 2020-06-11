package movie

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type SearchMoviesRequest struct {
	ApiKey string `json:"apikey"`
	Search string `json:"s"`
	Page   string `json:"page"`
}

type SearchMoviesRespon struct {
	Result interface{} `json:"Result"`
}

type Endpoint struct {
	SearchMoviesEndpoint endpoint.Endpoint
}

func MakeSearchMoviesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SearchMoviesRequest)
		movies, err := s.SearchMovies(ctx, req)
		return SearchMoviesRespon{Result: movies}, err
	}
}
