package repo

import (
	"context"
	"tiktok/internal/model"
	"tiktok/internal/repo/dao"
)

type UserRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	GetUserById(ctx context.Context, id int64) (*model.User, error)
}

type userRepo struct {
	dao dao.UserDao
}

func NewUserRepo(dao dao.UserDao) UserRepo {
	return &userRepo{
		dao: dao,
	}
}

func (u *userRepo) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	panic("implement me")
}

func (u *userRepo) CreateUser(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (u *userRepo) GetUserById(ctx context.Context, id int64) (*model.User, error) {
	panic("implement me")
}
