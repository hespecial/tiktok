package dao

import (
	"context"
	"gorm.io/gorm"
	"tiktok/internal/model"
)

type UserDao interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{
		db: db,
	}
}

func (u *userDao) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := u.db.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	return &user, err
}

func (u *userDao) CreateUser(ctx context.Context, user *model.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userDao) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	var user model.User
	return &user, u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
}
