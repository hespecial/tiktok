package repo

import (
	"context"
	"tiktok/internal/repo/cache"

	"tiktok/common/enum"
	"tiktok/internal/model"
	"tiktok/internal/repo/dao"
)

type FollowRepo interface {
	GetRelation(ctx context.Context, userId, followId int64) (*model.Follow, error)
	CreateRelation(ctx context.Context, userId, followId int64) error
	UpdateRelation(ctx context.Context, userId, followId int64, action enum.Relation) error
	GetFollowIds(ctx context.Context, userId int64, _type enum.Relation) ([]int64, error)
	GetFollowCount(ctx context.Context, userId int64, _type enum.Relation) (int64, error)
}

type followRepo struct {
	dao   dao.FollowDao
	cache cache.FollowCache
}

func NewFollowRepo(dao dao.FollowDao, cache cache.FollowCache) FollowRepo {
	return &followRepo{
		dao:   dao,
		cache: cache,
	}
}

func (f *followRepo) GetRelation(ctx context.Context, userId, followId int64) (*model.Follow, error) {
	return f.dao.GetRelation(ctx, userId, followId)
}

func (f *followRepo) CreateRelation(ctx context.Context, userId, followId int64) error {
	return f.dao.CreateRelation(ctx, userId, followId)
}

func (f *followRepo) UpdateRelation(ctx context.Context, userId, followId int64, action enum.Relation) error {
	return f.dao.UpdateRelation(ctx, userId, followId, action)
}

func (f *followRepo) GetFollowIds(ctx context.Context, userId int64, _type enum.Relation) ([]int64, error) {
	return f.dao.GetFollowIds(ctx, userId, _type)
}

func (f *followRepo) GetFollowCount(ctx context.Context, userId int64, _type enum.Relation) (int64, error) {
	return f.dao.GetFollowCount(ctx, userId, _type)
}
