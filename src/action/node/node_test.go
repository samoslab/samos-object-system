package node

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/node"
)

func TestAll(t *testing.T) {
	u := &node.Node{}
	u, r := node.Set(u)
	t.Log(u, r)
	u1, r1 := node.Get(u.NodeId)
	spew.Dump(u1, r1)
	list, err2 := node.List()
	spew.Dump(list, err2)
}
