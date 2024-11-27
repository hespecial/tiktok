package service

import (
	"tiktok/internal/repo/dao"
	"tiktok/util"
)

type UserService interface {
	GetUserByUsername(username string) *dao.User
	CreateUser(username, password string) error
}

type UserServiceImpl struct {
	repo dao.UserRepo
}

func NewUserService(repo dao.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) GetUserByUsername(username string) (*dao.User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *UserServiceImpl) CreateUser(username, password string) error {
	user := &dao.User{
		Username: username,
		Password: util.Encrypt(password),
	}
	return s.repo.CreateUser(user)
}
