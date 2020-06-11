package main

import (
	"log"
	"stockbit/gateway/movie"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()
	g.GET("/search", movie.GetSearchMovie)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
