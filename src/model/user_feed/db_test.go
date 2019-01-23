package user_feed

import (
	"testing"
)

func TestList(t *testing.T) {
	userId := "1533635654427489361"
	result, err := List(userId)
	t.Log(result, err)
}

func TestAdd(t *testing.T) {
	userId := "1533635654427489361"
	feedId := "1533635654427489361"
	err := Add(userId, feedId)
	t.Log(err)
}

func TestRemove(t *testing.T) {
	userId := "1533635654427489361"
	feedId := "1533635654427489361"
	err := Remove(userId, feedId)
	t.Log(err)
}
