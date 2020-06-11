package movie

import (
	"context"
)

//Service interface
type Service interface {
	SearchMovies(ctx context.Context, req SearchMoviesRequest) (interface{}, error)
}

//Repository interface
type Repository interface {
	GetListMoviesAPI(ctx context.Context, req SearchMoviesRequest) (interface{}, error)
}
