package user

import (
	"encoding/json"
	"time"

	"github.com/samoslab/samos-object-system/src/library/boltdb"
	"github.com/samoslab/samos-object-system/src/model/common"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

// 获取所有用户列表
func List() ([]*User, error) {
	list, err := boltdb.List(common.StoragePath, common.UserBucket)
	if err != nil {
		return nil, err
	}
	result := make([]*User, 0)
	for _, v := range list {
		u := &User{}
		err2 := json.Unmarshal(v, u)
		if err2 != nil {
			continue
		}
		result = append(result, u)
	}
	return result, nil
}

// 保存用户
func Set(u *User) (*User, error) {
	if u.UserId == "" {
		u.UserId = ticket.Generate()
	}
	if u.CreateTime == 0 {
		u.CreateTime = time.Now().Unix()
	}
	u.UpdateTime = time.Now().Unix()

	body, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	r := boltdb.Set(common.StoragePath, common.UserBucket, u.UserId, body)
	return u, r
}

// 获取单个用户信息
func Get(userId string) (*User, error) {
	result, err1 := boltdb.Get(common.StoragePath, common.UserBucket, userId)
	if err1 != nil {
		return nil, err1
	}
	u := &User{}
	err2 := json.Unmarshal([]byte(result), u)
	if err2 != nil {
		return nil, err2
	}
	return u, nil
}
