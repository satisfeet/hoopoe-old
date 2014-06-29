package router

import (
    "net/http"

    "github.com/julienschmidt/httprouter"
)

type Router struct {
    router    *httprouter.Router
    handles   []Handle
}

type Handle func(*Context)

func New() *Router {
    r := httprouter.New()

    return &Router{r, []Handle{}}
}

func (r *Router) Use(handle Handle) {
    r.handles = append(r.handles, handle)
}

func (r *Router) Get(pattern string, handle Handle) {
    r.router.Handle("GET", pattern, mediate(r, handle))
}

func (r *Router) Put(pattern string, handle Handle) {
    r.router.Handle("PUT", pattern, mediate(r, handle))
}

func (r *Router) Pos(pattern string, handle Handle) {
    r.router.Handle("POST", pattern, mediate(r, handle))
}

func (r *Router) Del(pattern string, handle Handle) {
    r.router.Handle("DELETE", pattern, mediate(r, handle))
}

func (r *Router) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
    r.router.ServeHTTP(writer, request)
}

func mediate(router *Router, handle Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        c := &Context{w, r, p, append(router.handles, handle), 0}

        c.Next()
    }
}
