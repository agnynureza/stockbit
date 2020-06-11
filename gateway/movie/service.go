package movie

import (
	"net/http"
	"stockbit/gateway/pb"

	"github.com/gin-gonic/gin"
)

func GetSearchMovie(ctx *gin.Context) {
	var apikey = ctx.Request.URL.Query()["apikey"]
	var pagination = ctx.Request.URL.Query()["pagination"]
	var searchword = ctx.Request.URL.Query()["searchword"]

	if len(apikey) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "No API key provided."})
		return
	}

	if len(searchword) < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Something went wrong."})
		return
	}

	if len(pagination) < 1 {
		pagination = append(pagination, "1")
	}

	req := &pb.Request{ApiKey: apikey[0], Search: searchword[0], Page: pagination[0]}
	result, err := ForwardRequest(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}
