package models

type User struct {
	Username string
	Password string
}

func (user *User) TableName() string {
	return "user"
}
