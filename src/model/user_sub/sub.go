package user_sub

import "github.com/samoslab/samos-object-system/src/model/feed"

// 订阅
func Subscribe(userId, feedId string) {

}

// 退订
func Unsubscribe(userId, feedId string) {

}

// 获取订阅列表
func GetList(userId string) []*feed.Feed {
	feedList := make([]*feed.Feed, 0)
	feedList = append(feedList, &feed.Feed{})
	return feedList
}
