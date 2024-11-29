package cache

type FollowCache interface {
}

type followCache struct{}

func NewFollowCache() FollowCache {
	return &followCache{}
}
