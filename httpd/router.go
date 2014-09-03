package httpd

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/satisfeet/go-context"
	"github.com/satisfeet/go-handler"
)

type Handler interface {
	ServeHTTP(*Context)
}

type Router struct {
	router *httprouter.Router
}

func NewRouter() *Router {
	return &Router{
		router: httprouter.New(),
	}
}

func (r *Router) Handle(method, pattern string, handler Handler) {
	r.router.Handle(method, pattern, func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		m := make(map[string]string, len(p))

		for _, p := range p {
			m[p.Key] = p.Value
		}

		handler.ServeHTTP(&Context{
			Context: &context.Context{
				Params:   m,
				Request:  req,
				Response: w,
			},
		})
	})
}

func (r *Router) HandleFunc(method, pattern string, handler HandlerFunc) {
	r.Handle(method, pattern, HandlerFunc(handler))
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if h, p, _ := r.router.Lookup(request.Method, request.URL.Path); h == nil {
		handler.NotFound(writer, request)
	} else {
		h(writer, request, p)
	}
}

type HandlerFunc func(*Context)

func (f HandlerFunc) ServeHTTP(c *Context) {
	f(c)
}
