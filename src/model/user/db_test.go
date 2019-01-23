package user

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

func TestSetUser(t *testing.T) {
	u := &User{}
	u.UserId = ticket.Generate()
	u.Name = "吴"
	u.SecretKey = "zzzz, secret key"
	r := Set(u)
	t.Log(r)
}

func TestGetUser(t *testing.T) {
	userId := "1533635654427489361"
	u, err := Get(userId)
	t.Log(u, err)
}

func TestGetUserList(t *testing.T) {
	u, _ := List()
	spew.Dump(u)
}
