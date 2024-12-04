package cache

import "github.com/redis/go-redis/v9"

type FollowCache interface {
}

type followCache struct {
	rdb *redis.Client
}

func NewFollowCache(rdb *redis.Client) FollowCache {
	return &followCache{
		rdb: rdb,
	}
}
