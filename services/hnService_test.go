package services

import (
	"context"
	"test-service-grpc/cache"
	pb "test-service-grpc/pb"

	"testing"

	"github.com/peterhellberg/hn"
	"github.com/stretchr/testify/require"
)

func TestGetTopStories(t *testing.T) {
	server := &Server{
		Cache:    cache.NewCache("3m"),
		HnClient: hn.DefaultClient,
	}
	ctx := context.Background()
	req := &pb.TopStoriesRequest{}
	res, _ := server.GetTopStories(ctx, req)
	require.NotNil(t, res)
	if len(res.Stories) < 10 {
		t.Log("test failed : array has not the good length")
	}
}
func TestWhois(t *testing.T) {
	server := &Server{
		Cache:    cache.NewCache("3m"),
		HnClient: hn.DefaultClient,
	}
	ctx := context.Background()
	req := &pb.WhoisRequest{User: "fra"}
	res, _ := server.Whois(ctx, req)
	require.NotNil(t, res)
	require.NotEmpty(t, res.Karma)
	require.NotEmpty(t, res.Nick)
	require.NotEmpty(t, res.About)
	require.NotEmpty(t, res.JoinedAt)
}
