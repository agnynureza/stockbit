package movie

import (
	"context"
	"encoding/json"
	"stockbit/movie-service/pb"

	"github.com/golang/protobuf/ptypes/any"
)

func EncodeSearchMoviesRes(_ context.Context, r interface{}) (interface{}, error) {
	res := r.(SearchMoviesRespon)
	serialized, err := json.Marshal(res.Result)
	if err != nil {
		return nil, err
	}
	return &pb.Response{
		Result: &any.Any{
			Value: serialized,
		},
	}, nil
}

func DecodeSearchMoviesReq(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.Request)
	return SearchMoviesRequest{req.ApiKey, req.Search, req.Page}, nil
}
