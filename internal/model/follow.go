package model

type Follow struct {
	Id       int64
	UserId   int64
	FollowId int64
	Status   uint8
}

func (*Follow) TableName() string {
	return "follows"
}
