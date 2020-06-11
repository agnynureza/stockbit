package movie

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
)

type repository struct {
	logger log.Logger
}

func NewRepo(log log.Logger) Repository {
	return &repository{
		logger: log,
	}
}

func (r *repository) GetListMoviesAPI(ctx context.Context, req SearchMoviesRequest) (*http.Response, error) {
	uri := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%s", req.ApiKey, req.Search, req.Page)
	request, errRequest := http.NewRequest(http.MethodGet, uri, nil)
	if errRequest != nil {
		return nil, errRequest
	}

	request = request.WithContext(ctx)
	request.Header.Set("Content-type", "Application/json")

	cli := &http.Client{Timeout: time.Second}
	resp, err := cli.Do(request)

	return resp, err
}
