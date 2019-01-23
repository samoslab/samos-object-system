package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/samoslab/samos-object-system/src/action/feed"
	"github.com/samoslab/samos-object-system/src/action/node"
	"github.com/samoslab/samos-object-system/src/action/object"
	"github.com/samoslab/samos-object-system/src/action/ticket"
	"github.com/samoslab/samos-object-system/src/action/user"
)

type Router struct {
	routerList map[string]func(w http.ResponseWriter, r *http.Request)
}

func New() *Router {
	router := &Router{}
	router.routerList = make(map[string]func(w http.ResponseWriter, r *http.Request))
	// feed
	router.routerList["/feed/get"] = feed.Get
	router.routerList["/feed/getlist"] = feed.GetList
	router.routerList["/feed/create"] = feed.Create
	// node
	router.routerList["/node/create"] = node.Create
	router.routerList["/node/get"] = node.Get
	// object
	router.routerList["/object/create"] = object.Create
	router.routerList["/object/delete"] = object.Delete
	router.routerList["/object/get"] = object.Get
	// feed
	router.routerList["/feed/create"] = feed.Create

	// user
	router.routerList["/user/create"] = user.Create
	router.routerList["/user/get"] = user.Get
	router.routerList["/user/feedlist"] = feed.Create

	// ticket
	router.routerList["/ticket/get"] = ticket.Get

	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.ToLower(r.URL.Path)
	f, err := router.routerList[path]
	if err != false {
		f(w, r)
	} else {
		fmt.Println("[no_router]", "\t", path)
	}
}
