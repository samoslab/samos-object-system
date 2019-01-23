package node

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

func TestSetNode(t *testing.T) {
	u := &Node{}
	u.NodeId = ticket.Generate()
	r := SaveNode(u)
	t.Log(r)
}

func TestGetNode(t *testing.T) {
	NodeId := "1533635654427489361"
	u, err := GetNode(NodeId)
	t.Log(u, err)
}

func TestGetNodeList(t *testing.T) {
	u, _ := GetNodeList()
	spew.Dump(u)
}
