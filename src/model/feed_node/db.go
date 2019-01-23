package feed_object

import (
	"github.com/samoslab/samos-object-system/src/library/boltdb"
	"github.com/samoslab/samos-object-system/src/model/common"
)

// 获取所有节点列表
func List(feedId string) ([]string, error) {
	bucket := common.FeedNodeBucket + ":" + feedId
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
func Add(feedId, nodeId string) error {
	bucket := common.FeedObjectBucket + ":" + feedId
	return boltdb.Set(common.StoragePath, bucket, nodeId, []byte(nodeId))
}

// 移除
func Remove(feedId, nodeId string) error {
	return nil
}
