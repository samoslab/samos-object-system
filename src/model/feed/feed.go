package feed

func Encode(o Feed) string {
	return ""
}

func Decode(s string) *Feed {
	o := &Feed{}
	return o
}

type Feed struct {
	FeedId     string
	NodeId     string
	Body       Body
	CreateTime uint32
	UpdateTime uint32
	NodeIdList []string
}
