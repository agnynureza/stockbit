package movie

import (
	"context"
	"net/http"
)

//Service interface
type Service interface {
	SearchMovies(ctx context.Context, req SearchMoviesRequest) (interface{}, error)
	GetMoviesListApiResponse(r *http.Response) (interface{}, error)
}

//Repository interface
type Repository interface {
	GetListMoviesAPI(ctx context.Context, req SearchMoviesRequest) (*http.Response, error)
}
