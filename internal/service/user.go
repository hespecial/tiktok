package service

import (
	"gorm.io/gorm"
	"tiktok/common/enum"
	"tiktok/internal/repo/dao"
	"tiktok/util"
)

type UserService interface {
	GetUserByUsername(username string) (*dao.User, error)
	CreateUser(username, password string) error

	GetUserInfo(id int64) (*UserInfo, error)
}

type UserServiceImpl struct {
	repo       dao.UserRepo
	followRepo dao.FollowRepo
}

func NewUserService(db *gorm.DB) UserService {
	return &UserServiceImpl{
		repo:       dao.NewUserDao(db),
		followRepo: dao.NewFollowDao(db),
	}
}

type UserInfo struct {
	Id             int64  `json:"id"`
	Username       string `json:"username"`
	FollowingCount int64  `json:"following_count"`
	FollowerCount  int64  `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
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

func (s *UserServiceImpl) GetUserInfo(id int64) (*UserInfo, error) {
	followingCnt, err := s.followRepo.GetFollowCount(id, enum.RelationFollowing)
	if err != nil {
		return nil, err
	}
	followerCnt, err := s.followRepo.GetFollowCount(id, enum.RelationFollower)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.GetUserById(id)
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
