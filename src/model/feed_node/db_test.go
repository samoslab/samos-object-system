package feed_object

import (
	"testing"
)

func TestList(t *testing.T) {
	feedId := "1533635654427489361"
	result, err := List(feedId)
	t.Log(result, err)
}

func TestAdd(t *testing.T) {
	feedId := "1533635654427489361"
	objectId := "1533635654427489361"
	err := Add(feedId, objectId)
	t.Log(err)
}

func TestRemove(t *testing.T) {
	feedId := "1533635654427489361"
	objectId := "1533635654427489361"
	err := Remove(feedId, objectId)
	t.Log(err)
}
