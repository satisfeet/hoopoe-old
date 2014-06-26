package mux

import (
    "net/http"
    "encoding/json"

    "github.com/julienschmidt/httprouter"
)

type Context struct {
    writer    http.ResponseWriter
    request   *http.Request
    params    httprouter.Params
    handles   []Handle
    index     int
}

func (c *Context) Get(header string) string {
    return c.request.Header.Get(header)
}

func (c *Context) Path() string {
    return c.request.URL.Path
}

func (c *Context) Method() string {
    return c.request.Method
}

func (c *Context) Params(p string) string {
    return c.params.ByName(p)
}

func (c *Context) Respond(b string, s int) {
    c.writer.WriteHeader(s)
    c.writer.Write([]byte(b))
}

func (c *Context) RespondJson(v interface{}, s int) {
    j, err := json.Marshal(v)

    if err != nil {
        c.RespondError(err, 500)

        return
    }

    c.writer.WriteHeader(s)
    c.writer.Write(j)
}

func (c *Context) RespondError(e error, s int) {
    c.Respond(e.Error(), s)
}

func (c *Context) Next() {
    c.index += 1

    c.handles[c.index - 1](c)
}
