package dao

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"tiktok/common/enum"
	"tiktok/internal/model"
)

type FollowDao interface {
	GetRelation(ctx context.Context, userId, followId int64) (*model.Follow, error)
	CreateRelation(ctx context.Context, userId, followId int64) error
	UpdateRelation(ctx context.Context, userId, followId int64, action enum.Relation) error
	GetFollowIds(ctx context.Context, userId int64, _type enum.Relation) ([]int64, error)
	GetFollowCount(ctx context.Context, userId int64, _type enum.Relation) (int64, error)
}

type followDao struct {
	db *gorm.DB
}

func NewFollowDao(db *gorm.DB) FollowDao {
	return &followDao{db: db}
}

func (f *followDao) GetRelation(ctx context.Context, userId, followId int64) (*model.Follow, error) {
	var relation model.Follow
	err := f.db.WithContext(ctx).
		Where("user_id = ? AND follow_id = ?", userId, followId).
		First(&relation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &relation, err
}

func (f *followDao) CreateRelation(ctx context.Context, userId, followId int64) error {
	return f.db.WithContext(ctx).
		Where("user_id = ? AND follow_id = ?", userId, followId).
		Create(&model.Follow{
			UserId:   userId,
			FollowId: followId,
			Status:   uint8(enum.RelationFollow),
		}).Error
}

func (f *followDao) UpdateRelation(ctx context.Context, userId, followId int64, action enum.Relation) error {
	var follow model.Follow
	return f.db.WithContext(ctx).
		Table(follow.TableName()).
		Where("user_id = ? AND follow_id = ?", userId, followId).
		Update("status", action).Error
}

func (f *followDao) GetFollowIds(ctx context.Context, userId int64, _type enum.Relation) ([]int64, error) {
	var list []int64
	var follow model.Follow
	idx := "user_id"
	if _type == enum.RelationFollower {
		idx = "follow_id"
	}
	return list, f.db.WithContext(ctx).
		Table(follow.TableName()).
		Where(idx+" = ? AND status = ?", userId, enum.RelationFollow).
		Pluck("follow_id", &list).Error
}

func (f *followDao) GetFollowCount(ctx context.Context, userId int64, _type enum.Relation) (int64, error) {
	var count int64
	var follow model.Follow
	idx := "user_id"
	if _type == enum.RelationFollower {
		idx = "follow_id"
	}
	return count, f.db.WithContext(ctx).
		Table(follow.TableName()).
		Where(idx+" = ? AND status = ?", userId, enum.RelationFollow).
		Count(&count).Error
}
