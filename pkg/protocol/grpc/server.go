package grpc

import (
	"context"
	"net"
	"os"
	"os/signal"

	"github.com/SrcHndWng/go-learning-gRPC-microservice/pkg/api/v1"
	"github.com/SrcHndWng/go-learning-gRPC-microservice/pkg/api/v2"
	"github.com/SrcHndWng/go-learning-gRPC-microservice/pkg/logger"
	"github.com/SrcHndWng/go-learning-gRPC-microservice/pkg/protocol/grpc/middleware"
	"google.golang.org/grpc"
)

// RunServer runs gRPC service to publish ToDo service
func RunServer(ctx context.Context, v1API v1.ToDoServiceServer, v2API v2.ToDoServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	// gRPC server statup options
	opts := []grpc.ServerOption{}

	// add middleware
	opts = middleware.AddLogging(logger.Log, opts)

	// register service
	server := grpc.NewServer(opts...)
	v1.RegisterToDoServiceServer(server, v1API)
	v2.RegisterToDoServiceServer(server, v2API)

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			logger.Log.Warn("shutting down gRPC server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	// start gRPC server
	logger.Log.Info("starting gRPC server...")
	return server.Serve(listen)
}
