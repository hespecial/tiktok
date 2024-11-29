package model

type User struct {
	Id       int64
	Username string
	Password string
}

func (*User) TableName() string {
	return "users"
}
