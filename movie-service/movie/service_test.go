package movie

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/stretchr/testify/assert"
)

var mockError = errors.New("error mock")

type M map[string]interface{}

func TestSearchMovie(t *testing.T) {
	result := make(M)
	search := make([]M, 0)

	s1 := M{
		"Title":  "Batman Begins",
		"Year":   "2005",
		"imdbID": "tt0372784",
		"Type":   "movie",
		"Poster": "https://m.media-amazon.com/images/M/MV5BZmUwNGU2ZmItMmRiNC00MjhlLTg5YWUtODMyNzkxODYzMmZlXkEyXkFqcGdeQXVyNTIzOTk5ODM@._V1_SX300.jpg",
	}

	s2 := M{
		"Title":  "Batman v Superman: Dawn of Justice",
		"Year":   "2016",
		"imdbID": "tt2975590",
		"Type":   "movie",
		"Poster": "https://m.media-amazon.com/images/M/MV5BYThjYzcyYzItNTVjNy00NDk0LTgwMWQtYjMwNmNlNWJhMzMyXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg",
	}
	search = append(search, s1, s2)

	result["search"] = search
	result["totalResults"] = "375"
	result["response"] = "True"
	req := SearchMoviesRequest{ApiKey: "faf7e5bb", Search: "Batman", Page: "1"}

	mockRepository := new(RepositoryMock)
	mockRepository.On("GetListMoviesAPI", context.Background(), req).Return(result, nil)
	service := NewService(mockRepository, nil)

	data, _ := service.SearchMovies(context.Background(), req)
	assert.Equal(t, data, result)
}

func TestSearchMovieError(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "movie",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	req := SearchMoviesRequest{ApiKey: "faf7e5bb", Search: "Batman", Page: "1"}

	mockRepository := new(RepositoryMock)
	mockRepository.On("GetListMoviesAPI", context.Background(), req).Return(nil, mockError)
	service := NewService(mockRepository, logger)

	_, err := service.SearchMovies(context.Background(), req)
	assert.Equal(t, err, mockError)
}
