package object

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/object"
)

func TestAll(t *testing.T) {
	o := &object.Object{}
	o.UserId = "zzzz"
	o.Body = "body"
	o, r := object.Set(o)
	t.Log(o, r)
	o1, r1 := object.Get(o.ObjectId)
	spew.Dump(o1, r1)
	list, err2 := object.List()
	spew.Dump(list, err2)
}
