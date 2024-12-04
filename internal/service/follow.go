package service

import (
	"context"
	"errors"
	"tiktok/common/enum"
	"tiktok/internal/repo"
)

type FollowService interface {
	Follow(ctx context.Context, userId, followId int64) error
	Unfollow(ctx context.Context, userId, followId int64) error
	GetFollowingList(ctx context.Context, userId int64) ([]*UserInfo, error)
	GetFollowerList(ctx context.Context, userId int64) ([]*UserInfo, error)
}

type FollowServiceImpl struct {
	repo     repo.FollowRepo
	userRepo repo.UserRepo
}

func NewFollowService(repo repo.FollowRepo, userRepo repo.UserRepo) FollowService {
	return &FollowServiceImpl{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (s *FollowServiceImpl) Follow(ctx context.Context, userId, followId int64) error {
	relation, err := s.repo.GetRelation(ctx, userId, followId)
	if err != nil {
		return err
	}
	if relation != nil {
		return s.repo.UpdateRelation(ctx, userId, followId, enum.RelationFollow)
	}
	return s.repo.CreateRelation(ctx, userId, followId)
}

func (s *FollowServiceImpl) Unfollow(ctx context.Context, userId, followId int64) error {
	relation, err := s.repo.GetRelation(ctx, userId, followId)
	if err != nil {
		return err
	}
	if relation == nil {
		return errors.New("bad request unfollow user")
	}
	return s.repo.UpdateRelation(ctx, userId, followId, enum.RelationUnfollow)
}

func (s *FollowServiceImpl) GetFollowingList(ctx context.Context, userId int64) ([]*UserInfo, error) {
	ids, err := s.repo.GetFollowIds(ctx, userId, enum.RelationFollowing)
	if err != nil {
		return nil, err
	}

	list := make([]*UserInfo, 0, len(ids))
	for _, id := range ids {
		user, err := s.userRepo.GetUserById(ctx, id)
		if err != nil {
			return nil, err
		}

		followingCnt, err := s.repo.GetFollowCount(ctx, id, enum.RelationFollowing)
		if err != nil {
			return nil, err
		}
		followerCnt, err := s.repo.GetFollowCount(ctx, id, enum.RelationFollower)
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

func (s *FollowServiceImpl) GetFollowerList(ctx context.Context, userId int64) ([]*UserInfo, error) {
	ids, err := s.repo.GetFollowIds(ctx, userId, enum.RelationFollower)
	if err != nil {
		return nil, err
	}

	list := make([]*UserInfo, 0, len(ids))
	for _, id := range ids {
		user, err := s.userRepo.GetUserById(ctx, id)
		if err != nil {
			return nil, err
		}

		followingCnt, err := s.repo.GetFollowCount(ctx, id, enum.RelationFollowing)
		if err != nil {
			return nil, err
		}
		followerCnt, err := s.repo.GetFollowCount(ctx, id, enum.RelationFollower)
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
		relation, err := s.repo.GetRelation(ctx, userId, id)
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
