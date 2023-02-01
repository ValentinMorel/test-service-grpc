package cache

import (
	"github.com/jellydator/ttlcache/v3"
	"time"
	pb "test-service-grpc/pb"
)

func NewCache(Expiration string) *ttlcache.Cache[string, []*pb.Story]{
	expiration, _ := time.ParseDuration(Expiration)
	cacheStories := ttlcache.New[string, []*pb.Story](
		ttlcache.WithTTL[string, []*pb.Story](expiration),
	)
	return cacheStories
}