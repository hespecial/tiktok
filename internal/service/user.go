package service

import "tiktok/internal/repo/dao"

type UserService interface{}

type UserServiceImpl struct {
	repo dao.UserRepo
}

func NewUserService(repo dao.UserRepo) *UserServiceImpl {
	return &UserServiceImpl{repo: repo}
}

func (s *UserServiceImpl) Login(username, password string) (interface{}, error) {
	panic("implement me")
}
