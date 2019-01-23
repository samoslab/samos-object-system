package user

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
	"github.com/samoslab/samos-object-system/src/model/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	t := "user/create"
	u := &user.User{}
	u.PublicKey = ""
	action.Output(w, t)
}
