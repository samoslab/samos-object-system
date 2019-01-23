package user

type User struct {
	UserId     string
	Name       string
	Email      string
	CreateTime int64
	UpdateTime int64
	PublicKey  string
}

// 获取feed列表
func (u *User) GetFeedCount() int {
	return 0
}
