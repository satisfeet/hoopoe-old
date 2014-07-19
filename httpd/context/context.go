package context

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Handler func(*Context)

type Context struct {
	// Param must be overwritten with a function
	// which returns a route parameter.
	Param func(string) string

	writer  http.ResponseWriter
	request *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer:  w,
		request: r,
	}
}

func (c *Context) Set(k, v string) {
	c.writer.Header().Set(k, v)
}

func (c *Context) Get(k string) string {
	return c.request.Header.Get(k)
}

func (c *Context) Query(k string) string {
	return c.request.URL.Query().Get(k)
}

func (c *Context) Error(err error, s int) {
	if err == nil {
		err = errors.New(http.StatusText(s))
	}

	// TODO: check if request accepts json
	http.Error(c.writer, `{"error":"`+err.Error()+`"}`, s)
}

func (c *Context) Parse(v interface{}) bool {
	var err error

	switch c.request.Header.Get("Content-Type") {
	case "application/json":
		err = json.NewDecoder(c.request.Body).Decode(v)
	}

	if err != nil {
		c.Error(err, http.StatusBadRequest)

		return false
	}

	return true
}

func (c *Context) Respond(v interface{}, s int) bool {
	var err error

	c.writer.WriteHeader(s)

	if v != nil {
		err = json.NewEncoder(c.writer).Encode(v)
	}

	return err == nil
}
