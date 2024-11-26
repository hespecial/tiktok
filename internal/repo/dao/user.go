package dao

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	Create(*User) error
}

func (*User) TableName() string {
	return "users"
}
