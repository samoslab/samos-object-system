package node

import "github.com/samoslab/samos-object-system/src/model/feed"

type Connection struct {
}

func (c *Connection) Subscribe() {

}

func (c *Connection) Unsubscribe() {

}

// 远程获取
func (c *Connection) Get(k string) *feed.Feed {
	f := &feed.Feed{}
	return f
}

// 主动推送，复制
func (c *Connection) Replicate(*feed.Feed) bool {
	return true
}
