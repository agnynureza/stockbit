package movie

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log/level"

	"github.com/go-kit/kit/log"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func NewService(rep Repository, log log.Logger) Service {
	return &service{
		repository: rep,
		logger:     log,
	}
}

func (s service) SearchMovies(ctx context.Context, req SearchMoviesRequest) (interface{}, error) {
	logger := log.With(s.logger, "method", "SearchMovies")
	resp, err := s.repository.GetListMoviesAPI(ctx, req)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	movies, err := s.GetMoviesListApiResponse(resp)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}
	return movies, nil
}

func (s service) GetMoviesListApiResponse(r *http.Response) (interface{}, error) {
	var response interface{}

	errDecode := json.NewDecoder(r.Body).Decode(&response)
	if errDecode != nil {
		return response, errDecode
	}
	return response, nil
}
