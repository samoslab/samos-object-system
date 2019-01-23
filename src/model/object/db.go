package object

import (
	"encoding/json"
	"time"

	"github.com/samoslab/samos-object-system/src/library/boltdb"
	"github.com/samoslab/samos-object-system/src/model/common"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

// 获取所有用户列表
func List() ([]*Object, error) {
	list, err := boltdb.List(common.StoragePath, common.ObjectBucket)
	if err != nil {
		return nil, err
	}
	result := make([]*Object, 0)
	for _, v := range list {
		u := &Object{}
		err2 := json.Unmarshal(v, u)
		if err2 != nil {
			continue
		}
		result = append(result, u)
	}
	return result, nil
}

// 保存
func Set(o *Object) (*Object, error) {
	if o.ObjectId == "" {
		o.ObjectId = ticket.Generate()
	}

	if o.CreateTime == 0 {
		o.CreateTime = time.Now().Unix()
	}
	o.UpdateTime = time.Now().Unix()
	body, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	err1 := boltdb.Set(common.StoragePath, common.ObjectBucket, o.ObjectId, body)
	return o, err1
}

// 获取
func Get(userId string) (*Object, error) {
	result, err1 := boltdb.Get(common.StoragePath, common.ObjectBucket, userId)
	if err1 != nil {
		return nil, err1
	}
	u := &Object{}
	err2 := json.Unmarshal([]byte(result), u)
	if err2 != nil {
		return nil, err2
	}
	return u, nil
}
