package movie

import (
	"encoding/json"
	"stockbit/gateway/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func ForwardRequest(ctx *gin.Context, req *pb.Request) (interface{}, error) {
	var result interface{}

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := pb.NewMoviesServiceClient(conn)

	resp, err := client.SearchMovies(ctx, req)
	if err := json.Unmarshal(resp.Result.Value, &result); err != nil {
		return nil, err
	}
	return result, nil
}
