package feed

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/feed"
)

func TestAll(t *testing.T) {
	u := &feed.Feed{}
	u, r := feed.Set(u)
	t.Log(u, r)
	u1, r1 := feed.Get(u.FeedId)
	spew.Dump(u1, r1)
	list, err2 := feed.List()
	spew.Dump(list, err2)
}
