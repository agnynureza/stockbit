package movie

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
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

	return resp, nil
}
