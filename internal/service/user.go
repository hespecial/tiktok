package service

import (
	"context"
	"tiktok/common/enum"
	"tiktok/internal/model"
	"tiktok/internal/repo"
	"tiktok/util"
)

type UserService interface {
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	CreateUser(ctx context.Context, username, password string) error

	GetUserInfo(ctx context.Context, id int64) (*UserInfo, error)
}

type UserServiceImpl struct {
	repo       repo.UserRepo
	followRepo repo.FollowRepo
}

func NewUserService(repo repo.UserRepo, followRepo repo.FollowRepo) UserService {
	return &UserServiceImpl{
		repo:       repo,
		followRepo: followRepo,
	}
}

type UserInfo struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	FollowingCount int64  `json:"following_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
}

func (s *UserServiceImpl) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.repo.GetUserByUsername(ctx, username)
}

func (s *UserServiceImpl) CreateUser(ctx context.Context, username, password string) error {
	user := &model.User{
		Username: username,
		Password: util.Encrypt(password),
	}
	return s.repo.CreateUser(ctx, user)
}

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, id int64) (*UserInfo, error) {
	followingCnt, err := s.followRepo.GetFollowCount(ctx, id, enum.RelationFollowing)
	if err != nil {
		return nil, err
	}
	followerCnt, err := s.followRepo.GetFollowCount(ctx, id, enum.RelationFollower)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	return &UserInfo{
		Id:             id,
		Username:       user.Username,
		FollowingCount: followingCnt,
		FollowerCount:  followerCnt,
		IsFollow:       false,
	}, err
}
