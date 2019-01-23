package node

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/samoslab/samos-object-system/src/library/boltdb"
	"github.com/samoslab/samos-object-system/src/model/common"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

// 获取所有节点列表
func List() ([]*Node, error) {
	list, err := boltdb.List(common.StoragePath, common.NodeBucket)
	if err != nil {
		return nil, err
	}
	result := make([]*Node, 0)
	for _, v := range list {
		u := &Node{}
		err2 := json.Unmarshal(v, u)
		if err2 != nil {
			continue
		}
		result = append(result, u)
	}
	return result, nil
}

// 保存节点
func Set(n *Node) (*Node, error) {
	if n.NodeId == "" {
		n.NodeId = ticket.Generate()
	}
	if n.CreateTime == 0 {
		n.CreateTime = time.Now().Unix()
	}
	n.UpdateTime = time.Now().Unix()
	body, err := json.Marshal(n)
	if err != nil {
		return nil, errors.New("encode failed")
	}
	r := boltdb.Set(common.StoragePath, common.NodeBucket, n.NodeId, body)
	return n, r
}

// 获取单个节点信息
func Get(NodeId string) (*Node, error) {
	result, err1 := boltdb.Get(common.StoragePath, common.NodeBucket, NodeId)
	if err1 != nil {
		return nil, err1
	}
	u := &Node{}
	err2 := json.Unmarshal([]byte(result), u)
	if err2 != nil {
		return nil, err2
	}
	return u, nil
}
