package services

import (
	"context"
	"time"
	"fmt"

	pb "test-service-grpc/pb"

	"github.com/peterhellberg/hn"
	"github.com/jellydator/ttlcache/v3"
	"errors"
)



// Implement interface to be registered on grpc server
type Server struct {
	HnClient *hn.Client
	Cache *ttlcache.Cache[string, []*pb.Story]
}

func (s *Server) GetTopStories(ctx context.Context, req *pb.TopStoriesRequest) (*pb.TopStories, error) {
	// Check cache if we have data ready to be sent
	if cacheItem := s.Cache.Get("latest"); cacheItem != nil {
		return &pb.TopStories{
			Stories: cacheItem.Value(),
		}, nil
	}
	
	// Declare a buffered array to avoid memory re-allocation with append
	stories := make([]*pb.Story, 10)
	ids, err := s.HnClient.TopStories()
	if err != nil {
		return &pb.TopStories{}, errors.New("Encountered an unexpected behavior from the hacker news API, TopStories method")
	}

	for i, id := range ids[:10] {
		item, err := s.HnClient.Item(id)
		if err != nil {
			panic(err)
		}
		// Assign to buffered array
		stories[i] = &pb.Story{
			Title: item.Title,
			Url:   item.URL,
		}
	}
	// Set newest data to cache for other calls coming too quickly
	s.Cache.Set("latest", stories, ttlcache.DefaultTTL)
	return &pb.TopStories{
		Stories: stories,
	}, nil
}

func (s *Server) Whois(ctx context.Context, in *pb.WhoisRequest) (*pb.User, error) {
	
	u, err := hn.DefaultClient.User(in.User)
	if err != nil {
		return &pb.User{}, errors.New("Encountered an unexpected behavior from the hacker news API, User method")
	}
	year, month, day := time.Unix(int64(u.Created), 0).Date()
	return &pb.User{
		Nick:     u.ID,
		Karma:    uint64(u.Karma),
		About:    u.About,
		JoinedAt: fmt.Sprintf("%d-%d-%d", year, month, day),
	}, nil

}
