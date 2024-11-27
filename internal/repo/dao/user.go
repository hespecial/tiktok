package dao

import "gorm.io/gorm"

type User struct {
	Id       int64
	Username string
	Password string
}

type UserRepo interface {
	GetUserByUsername(username string) (*User, error)
	CreateUser(user *User) error
}

type UserDao struct {
	db *gorm.DB
}

func (*User) TableName() string {
	return "users"
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{
		db: db,
	}
}

func (u *UserDao) GetUserByUsername(username string) (*User, error) {
	var user User
	err := u.db.Where("username = ?", username).Find(&user).Error
	return &user, err
}

func (u *UserDao) CreateUser(user *User) error {
	return u.db.Create(user).Error
}
