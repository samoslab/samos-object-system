package object

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

func TestSetObject(t *testing.T) {
	u := &Object{}
	u.ObjectId = ticket.Generate()
	r := Set(u)
	t.Log(r)
}

func TestGetObject(t *testing.T) {
	userId := "1533635654427489361"
	u, err := Get(userId)
	t.Log(u, err)
}

func TestGetObjectList(t *testing.T) {
	u, _ := List()
	spew.Dump(u)
}
