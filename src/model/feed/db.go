package feed

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/samoslab/samos-object-system/src/library/boltdb"
	"github.com/samoslab/samos-object-system/src/model/common"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

// 获取所有节点列表
func List() ([]*Feed, error) {
	list, err := boltdb.List(common.StoragePath, common.FeedBucket)
	if err != nil {
		return nil, err
	}
	result := make([]*Feed, 0)
	for _, v := range list {
		u := &Feed{}
		err2 := json.Unmarshal(v, u)
		if err2 != nil {
			continue
		}
		result = append(result, u)
	}
	return result, nil
}

// 保存节点
func Set(f *Feed) (*Feed, error) {
	if f.FeedId == "" {
		f.FeedId = ticket.Generate()
	}
	if f.CreateTime == 0 {
		f.CreateTime = time.Now().Unix()
	}
	f.UpdateTime = time.Now().Unix()
	body, err := json.Marshal(f)
	if err != nil {
		return nil, errors.New("encode failed")
	}
	r := boltdb.Set(common.StoragePath, common.FeedBucket, f.FeedId, body)
	return f, r
}

// 获取单个节点信息
func Get(FeedId string) (*Feed, error) {
	result, err1 := boltdb.Get(common.StoragePath, common.FeedBucket, FeedId)
	if err1 != nil {

	}
	u := &Feed{}
	err2 := json.Unmarshal([]byte(result), u)
	if err2 != nil {

	}
	return u, nil
}
