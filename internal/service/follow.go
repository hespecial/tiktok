package service

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/common/enum"
	"tiktok/internal/repo/dao"
)

type FollowService interface {
	Follow(userId, followId int64) error
	Unfollow(userId, followId int64) error
	GetFollowingList(userId int64) ([]*UserInfo, error)
	GetFollowerList(userId int64) ([]*UserInfo, error)
}

type FollowServiceImpl struct {
	repo     *dao.FollowDao
	userRepo *dao.UserDao
}

func NewFollowService(db *gorm.DB) FollowService {
	return &FollowServiceImpl{
		repo:     dao.NewFollowDao(db),
		userRepo: dao.NewUserDao(db),
	}
}

func (s *FollowServiceImpl) Follow(userId, followId int64) error {
	relation, err := s.repo.GetRelation(userId, followId)
	if err != nil {
		return err
	}
	if relation != nil {
		return s.repo.UpdateRelation(userId, followId, enum.RelationFollow)
	}
	return s.repo.CreateRelation(userId, followId)
}

func (s *FollowServiceImpl) Unfollow(userId, followId int64) error {
	relation, err := s.repo.GetRelation(userId, followId)
	if err != nil {
		return err
	}
	if relation == nil {
		return errors.New("bad request unfollow user")
	}
	return s.repo.UpdateRelation(userId, followId, enum.RelationUnfollow)
}

func (s *FollowServiceImpl) GetFollowingList(userId int64) ([]*UserInfo, error) {
	ids, err := s.repo.GetFollowIds(userId, enum.RelationFollowing)
	if err != nil {
		return nil, err
	}

	var list []*UserInfo
	for _, id := range ids {
		user, err := s.userRepo.GetUserById(id)
		if err != nil {
			return nil, err
		}

		followingCnt, err := s.repo.GetFollowCount(id, enum.RelationFollowing)
		if err != nil {
			return nil, err
		}
		followerCnt, err := s.repo.GetFollowCount(id, enum.RelationFollower)
		if err != nil {
			return nil, err
		}

		userInfo := &UserInfo{
			Id:             user.Id,
			Username:       user.Username,
			FollowingCount: followingCnt,
			FollowerCount:  followerCnt,
			IsFollow:       true, // 当前用户一定关注了其关注列表的所有用户
		}

		list = append(list, userInfo)
	}

	return list, nil
}

func (s *FollowServiceImpl) GetFollowerList(userId int64) ([]*UserInfo, error) {
	ids, err := s.repo.GetFollowIds(userId, enum.RelationFollower)
	if err != nil {
		return nil, err
	}

	var list []*UserInfo
	for _, id := range ids {
		user, err := s.userRepo.GetUserById(id)
		if err != nil {
			return nil, err
		}

		followingCnt, err := s.repo.GetFollowCount(id, enum.RelationFollowing)
		if err != nil {
			return nil, err
		}
		followerCnt, err := s.repo.GetFollowCount(id, enum.RelationFollower)
		if err != nil {
			return nil, err
		}

		userInfo := &UserInfo{
			Id:             user.Id,
			Username:       user.Username,
			FollowingCount: followingCnt,
			FollowerCount:  followerCnt,
			IsFollow:       false,
		}

		// 用户是否关注了他的粉丝
		relation, err := s.repo.GetRelation(userId, id)
		if err != nil {
			return nil, err
		}
		if relation != nil && relation.Status == uint8(enum.RelationFollow) {
			userInfo.IsFollow = true
		}

		list = append(list, userInfo)
	}
	return list, nil
}
