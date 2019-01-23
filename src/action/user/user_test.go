package user

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/user"
)

func TestAll(t *testing.T) {
	u := &user.User{}
	u.PublicKey = "bbb"
	u, r := user.Set(u)
	t.Log(u, r)
	u1, r1 := user.Get(u.UserId)
	spew.Dump(u1, r1)
	list, err2 := user.List()
	spew.Dump(list, err2)
}
