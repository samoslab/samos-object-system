package node

import "github.com/samoslab/samos-object-system/src/model/object"

func New() *Node {
	n := &Node{}
	n.Port = "9"
	return n
}

type Node struct {
	NodeId     string `名称`
	NodeName   string
	Ip         string
	Port       string
	CreateTime int64
	UpdateTime int64
}

// 注册发现
func (n *Node) Register() {

}

// 远程获取
func (n *Node) GetObject(objectId string) *object.Object {
	f := &object.Object{}
	return f
}

// 主动推送，复制
func (n *Node) ReplicateObject(object *object.Object) bool {
	return true
}

// 注册发现
func (n *Node) getServerIp() string {
	return "127.0.0.1"
}
