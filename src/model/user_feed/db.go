package user_feed

import (
	"github.com/samoslab/samos-object-system/src/library/boltdb"
	"github.com/samoslab/samos-object-system/src/model/common"
)

// 获取列表
func List(userId string) ([]string, error) {
	bucket := common.UserFeedBucket + ":" + userId
	list, err := boltdb.List(common.StoragePath, bucket)
	if err != nil {
		return nil, err
	}
	result := make([]string, 0)
	for k, _ := range list {
		result = append(result, k)
	}
	return result, nil
}

// 添加
func Add(userId, feedId string) error {
	bucket := common.UserFeedBucket + ":" + userId
	return boltdb.Set(common.StoragePath, bucket, feedId, []byte(feedId))
}

// 移除
func Remove(userId, feedId string) error {
	return nil
}
