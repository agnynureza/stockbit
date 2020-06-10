package main

import (
	"encoding/json"
	"log"
	"net/http"
	"stockbit/gateway/pb"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewMoviesServiceClient(conn)
	g := gin.Default()
	g.GET("/search", func(ctx *gin.Context) {
		var apikey = ctx.Request.URL.Query()["apikey"]
		var pagination = ctx.Request.URL.Query()["pagination"]
		var searchword = ctx.Request.URL.Query()["searchword"]

		if len(apikey) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "url params apikey is required"})
			return
		}

		if len(searchword) < 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "url params searchword is required"})
			return
		}

		if len(pagination) < 1 {
			pagination = append(pagination, "1")
		}

		req := &pb.Request{ApiKey: apikey[0], Search: searchword[0], Page: pagination[0]}
		if response, err := client.SearchMovies(ctx, req); err == nil {
			var resp interface{}
			if err := json.Unmarshal(response.Result.Value, &resp); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, gin.H{
				"Result": resp,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
