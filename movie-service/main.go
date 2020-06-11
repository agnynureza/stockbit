package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"stockbit/movie-service/movie"
	"stockbit/movie-service/pb"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"
)

func main() {
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

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	ctx := context.Background()
	var srv movie.Service
	{
		repository := movie.NewRepo(logger)

		srv = movie.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		listener, err := net.Listen("tcp", ":9090")
		if err != nil {
			errs <- err
			return
		}
		gRPCServer := grpc.NewServer()
		pb.RegisterMoviesServiceServer(gRPCServer, movie.NewGRPCServer(ctx, movie.Endpoint{
			SearchMoviesEndpoint: movie.MakeSearchMoviesEndpoint(srv),
		}))

		fmt.Println("gRPC listen on 9090")
		errs <- gRPCServer.Serve(listener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
