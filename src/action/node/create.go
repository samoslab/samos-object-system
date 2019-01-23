package node

import (
	"net/http"

	"github.com/samoslab/samos-object-system/src/action"
	"github.com/samoslab/samos-object-system/src/model/node"
)

func Create(w http.ResponseWriter, r *http.Request) {
	ip := r.FormValue("Ip")
	n := &node.Node{}
	n.NodeId = ip
	n, _ = node.Set(n)
	action.Output(w, "node/create", ip, n)
}
