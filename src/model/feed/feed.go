package feed

import "github.com/samoslab/samos-object-system/src/model/cell"

func Encode(o Feed) string {
	return ""
}

func Decode(s string) *Feed {
	o := &Feed{}
	return o
}

type Feed struct {
	FeedId      string
	NodeId      string
	Version     string
	VersionList []string
	cellId      string
	Cell        cell.Cell
	CreateTime  uint32
	UpdateTime  uint32
	NodeIdList  []string
}
