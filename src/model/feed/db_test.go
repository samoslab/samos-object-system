package feed

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

func TestSetFeed(t *testing.T) {
	u := &Feed{}
	u.FeedId = ticket.Generate()
	r := Set(u)
	t.Log(r)
}

func TestGetFeed(t *testing.T) {
	FeedId := "1533635654427489361"
	u, err := Get(FeedId)
	t.Log(u, err)
}

func TestGetFeedList(t *testing.T) {
	u, _ := List()
	spew.Dump(u)
}
