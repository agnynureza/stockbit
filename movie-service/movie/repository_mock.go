package movie

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m RepositoryMock) GetListMoviesAPI(ctx context.Context, req SearchMoviesRequest) (resp interface{}, err error) {
	args := m.Called(ctx, req)
	data, ok := args.Get(0).(interface{})
	if !ok {
		return resp, args.Error(1)
	}

	return data, args.Error(1)
}
