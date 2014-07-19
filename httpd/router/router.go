package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satisfeet/hoopoe/httpd/context"
)

const (
	MethodShow   = "GET"
	MethodCreate = "POST"
	MethodUpdate = "PUT"
	MethodDelete = "DELETE"
)

type Router struct {
	router *httprouter.Router
}

func NewRouter() *Router {
	r := httprouter.New()

	return &Router{
		router: r,
	}
}

func (router *Router) Handle(m, p string, h Handler) {
	router.router.Handle(m, p, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c := context.NewContext(w, r)
		c.Param = p.ByName

		h.ServeHTTP(c)
	})
}

func (router *Router) HandleFunc(m, p string, h HandlerFunc) {
	router.Handle(m, p, HandlerFunc(h))
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
