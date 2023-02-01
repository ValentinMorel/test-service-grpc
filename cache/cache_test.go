package cache

import (
	"time"
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/jellydator/ttlcache/v3"
	pb "test-service-grpc/pb"
)

func TestNewCache(t *testing.T) {
	c := ttlcache.New[string, []*pb.Story](
		ttlcache.WithTTL[string, []*pb.Story](time.Minute),
		ttlcache.WithCapacity[string, []*pb.Story](1),
	)
	require.NotNil(t, c)
}
