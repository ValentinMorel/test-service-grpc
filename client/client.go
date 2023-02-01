package client

import (
	pb "test-service-grpc/pb"

	"google.golang.org/grpc"
)

func Connect(address string) pb.HnServiceClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := pb.NewHnServiceClient(conn)
	return client
}
