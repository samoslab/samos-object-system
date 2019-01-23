package object

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
	"github.com/samoslab/samos-object-system/src/model/object"
)

func Get(w http.ResponseWriter, r *http.Request) {
	FeedId := r.FormValue("FeedId")
	UserId := r.FormValue("UserId")
	Body := r.FormValue("Body")

	o := &object.Object{}
	o.FeedId = FeedId
	o.UserId = UserId
	o.Body = Body
	o, _ = object.Set(o)
	action.Output(w, "object/create", o)
}
