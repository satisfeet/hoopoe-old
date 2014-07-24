package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satisfeet/hoopoe/httpd/context"
)

// Constants to represent HTTP methods as CRUD actions.
const (
	MethodShow   = "GET"
	MethodCreate = "POST"
	MethodUpdate = "PUT"
	MethodDelete = "DELETE"
)

// Router wraps httprouter to be used with custom http handlers.
type Router struct {
	router *httprouter.Router
}

// Returns an initialized router.
func NewRouter() *Router {
	r := httprouter.New()

	return &Router{
		router: r,
	}
}

// Handler is an interface which is similar to http.Handler except that the
// function takes a context object over a response and request.
type Handler interface {
	ServeHTTP(*context.Context)
}

// Same as Handler just for standalone functions.
type HandlerFunc func(*context.Context)

// Compability to convert standalone functions to Handler interface.
func (handler HandlerFunc) ServeHTTP(c *context.Context) {
	handler(c)
}

// Adds a Handler to the router and wraps the provided arguments in a Context.
func (router *Router) Handle(m, p string, h Handler) {
	router.router.Handle(m, p, func(w http.ResponseWriter, r *http.Request,
		p httprouter.Params) {
		c := context.NewContext(w, r)
		c.Param = p.ByName

		h.ServeHTTP(c)
	})
}

// Adds a HandlerFunc to the router.
func (router *Router) HandleFunc(m, p string, h HandlerFunc) {
	router.Handle(m, p, HandlerFunc(h))
}

// Forwards requests to the internal router.
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.router.ServeHTTP(w, r)
}
