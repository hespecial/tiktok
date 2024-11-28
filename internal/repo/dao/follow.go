package dao

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/common/enum"
)

type Follow struct {
	Id       int64
	UserId   int64
	FollowId int64
	Status   uint8
}

func (*Follow) TableName() string {
	return "follows"
}

type FollowRepo interface {
	GetRelation(userId, followId int64) (*Follow, error)
	CreateRelation(userId, followId int64) error
	UpdateRelation(userId, followId int64, action enum.Relation) error
	GetFollowIds(userId int64, _type enum.Relation) ([]int64, error)
	GetFollowCount(userId int64, _type enum.Relation) (int64, error)
}

type FollowDao struct {
	db *gorm.DB
}

func NewFollowDao(db *gorm.DB) *FollowDao {
	return &FollowDao{db: db}
}

func (f *FollowDao) GetRelation(userId, followId int64) (*Follow, error) {
	var relation Follow
	err := f.db.Where("user_id = ? AND follow_id = ?", userId, followId).
		First(&relation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &relation, err
}

func (f *FollowDao) CreateRelation(userId, followId int64) error {
	return f.db.Where("user_id = ? AND follow_id = ?", userId, followId).
		Create(&Follow{
			UserId:   userId,
			FollowId: followId,
			Status:   uint8(enum.RelationFollow),
		}).Error
}

func (f *FollowDao) UpdateRelation(userId, followId int64, action enum.Relation) error {
	var follow Follow
	return f.db.Table(follow.TableName()).
		Where("user_id = ? AND follow_id = ?", userId, followId).
		Update("status", action).Error
}

func (f *FollowDao) GetFollowIds(userId int64, _type enum.Relation) ([]int64, error) {
	var list []int64
	var follow Follow
	idx := "user_id"
	if _type == enum.RelationFollower {
		idx = "follow_id"
	}
	return list, f.db.Table(follow.TableName()).
		Where(idx+" = ? AND status = ?", userId, enum.RelationFollow).
		Pluck("follow_id", &list).Error
}

func (f *FollowDao) GetFollowCount(userId int64, _type enum.Relation) (int64, error) {
	var count int64
	var follow Follow
	idx := "user_id"
	if _type == enum.RelationFollower {
		idx = "follow_id"
	}
	return count, f.db.Table(follow.TableName()).
		Where(idx+" = ? AND status = ?", userId, enum.RelationFollow).
		Count(&count).Error
}
