package movie

import (
	"context"
	"stockbit/movie-service/pb"

	gt "github.com/go-kit/kit/transport/grpc"
)

type GRPCServer struct {
	searchMovies gt.Handler
}

func (s *GRPCServer) SearchMovies(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	_, resp, err := s.searchMovies.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Response), nil
}

func NewGRPCServer(_ context.Context, endpoint Endpoint) pb.MoviesServiceServer {
	return &GRPCServer{
		searchMovies: gt.NewServer(
			endpoint.SearchMoviesEndpoint,
			DecodeSearchMoviesReq,
			EncodeSearchMoviesRes,
		),
	}
}
