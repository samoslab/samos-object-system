package feed

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
)

func Get(w http.ResponseWriter, r *http.Request) {
	feedId := r.FormValue("FeedId")
	action.Output(w, "feed/get", feedId)
}
