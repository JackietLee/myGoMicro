package main

import (
	"first/handler"
	pb "first/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "first"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterFirstHandler(srv.Server(), new(handler.First)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
