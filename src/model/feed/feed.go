package feed

type Feed struct {
	FeedId       string
	UserId       string
	Type         uint8
	LastObjectId string
	CreateTime   int64
	UpdateTime   int64
}

func (f Feed) GetObjectCount() int {
	return 0
}
