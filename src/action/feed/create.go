package feed

import (
	"encoding/json"
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
	"github.com/samoslab/samos-object-system/src/model/feed"
)

func Create(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("UserId")
	f := &feed.Feed{}
	f.UserId = userId
	f, _ = feed.Set(f)
	result, _ := json.Marshal(f)
	action.Output(w, string(result))
}
