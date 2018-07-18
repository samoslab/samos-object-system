package cell

import "github.com/samoslab/samos-object-system/src/model/feed"

// 最小存储单元
type Cell struct {
	CellId     string
	CellName   string
	Address    string
	Body       string
	FeedId     string
	Feed       feed.Feed ``
	CreateTime string
}
