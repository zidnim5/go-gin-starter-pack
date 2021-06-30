package entities

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (user *User) TableName() string {
	return "user"
}
