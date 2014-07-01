package router

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	writer  http.ResponseWriter
	request *http.Request
	params  httprouter.Params
	handles []Handle
	index   int
}

func (c *Context) Get(header string) string {
	return c.request.Header.Get(header)
}

func (c *Context) Path() string {
	return c.request.URL.Path
}

func (c *Context) Param(param string) string {
	return c.params.ByName(param)
}

func (c *Context) Query() url.Values {
	return c.request.URL.Query()
}

func (c *Context) Method() string {
	return c.request.Method
}

func (c *Context) ParseJson(value interface{}) error {
	return json.NewDecoder(c.request.Body).Decode(value)
}

func (c *Context) RespondJson(value interface{}, status int) {
	j, err := json.Marshal(value)

	if err != nil {
		c.RespondError(err, 500)

		return
	}

	c.writer.WriteHeader(status)
	c.writer.Write(j)
}

func (c *Context) RespondError(err error, status int) {
	if err == nil {
		err = errors.New(http.StatusText(status))
	}

	c.Respond(err.Error(), status)
}

func (c *Context) Respond(body string, status int) {
	c.writer.WriteHeader(status)
	c.writer.Write([]byte(body))
}

func (c *Context) Next() {
	c.index += 1

	c.handles[c.index-1](c)
}
