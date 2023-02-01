package server

import (
	"log"
	"net"
	"strconv"

	"test-service-grpc/config"
	pb "test-service-grpc/pb"

	"test-service-grpc/services"

	"test-service-grpc/cache"

	"github.com/peterhellberg/hn"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start(cfg *config.Config) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// define the grpcServer that will listen to http connection
	grpcServer := grpc.NewServer()

	// define the server struct with reimplemented grpc methods
	server := &services.Server{
		Cache:    cache.NewCache(cfg.CacheExpiration),
		HnClient: hn.DefaultClient,
	}
	pb.RegisterHnServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	go server.Cache.Start()

	log.Printf("Started server on :%d", cfg.Port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
