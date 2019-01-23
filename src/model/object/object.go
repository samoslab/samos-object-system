package object

type Object struct {
	ObjectId     string
	FeedId       string
	UserId       string
	Body         string
	CreateNodeId string
	CreateTime   int64
	UpdateTime   int64
}

func Encode(o Object) string {
	return ""
}

func Decode(s string) *Object {
	o := &Object{}
	return o
}
