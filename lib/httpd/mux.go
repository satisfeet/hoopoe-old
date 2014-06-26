package httpd

import (
  "net/http"

  "github.com/julienschmidt/httprouter"
)

type Mux struct {
  router    *httprouter.Router
  handles   []Handle
}

type Handle func(*Context)

func NewMux() *Mux {
  r := httprouter.New()

  return &Mux{r, []Handle{}}
}

func (m *Mux) Use(h Handle) {
  m.handles = append(m.handles, h)
}

func (m *Mux) Get(p string, h Handle) {
  m.router.GET(p, mediate(m, h))
}

func (m *Mux) Put(p string, h Handle) {
  m.router.PUT(p, mediate(m, h))
}

func (m *Mux) Post(p string, h Handle) {
  m.router.POST(p, mediate(m, h))
}

func (m *Mux) Delete(p string, h Handle) {
  m.router.DELETE(p, mediate(m, h))
}

func (m *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  m.router.ServeHTTP(w, r)
}

func mediate(m *Mux, h Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    c := &Context{w, r, p, append(m.handles, h), 0}

    c.Next()
  }
}
