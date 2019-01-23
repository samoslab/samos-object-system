package feed

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("UserId")
	action.Output(w, "feed/getlist", userId)
}
