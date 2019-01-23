package ticket

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
	"github.com/samoslab/samos-object-system/src/model/ticket"
)

func Get(w http.ResponseWriter, r *http.Request) {
	t := ticket.Generate()
	action.Output(w, t)
}
