package node

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
)

func Get(w http.ResponseWriter, r *http.Request) {
	nodeId := r.FormValue("NodeId")
	action.Output(w, "node/get", nodeId)
}
