package context

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Context mediates to a http request response pair but adds shortcuts and
// convenient handler for use in REST APIs.
type Context struct {
	// Param must be overwritten with a function
	// which returns a route parameter.
	Param func(string) string

	writer  http.ResponseWriter
	request *http.Request
}

// Returns an initialized Context.
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		writer:  w,
		request: r,
	}
}

// Sets the header of a response.
func (c *Context) Set(k, v string) {
	c.writer.Header().Set(k, v)
}

// Returns the header of a request.
func (c *Context) Get(k string) string {
	return c.request.Header.Get(k)
}

// Returns a query string value by key.
func (c *Context) Query(k string) string {
	return c.request.URL.Query().Get(k)
}

// Responds an json encoded error. If no error is provided will use the standard
// http status text for the given error code.
func (c *Context) Error(err error, s int) {
	if err == nil {
		err = errors.New(http.StatusText(s))
	}

	// TODO: check if request accepts json
	http.Error(c.writer, `{"error":"`+err.Error()+`"}`, s)
}

// Parses a request body and maps the data to the provided value. If an error
// occurs it will respond an error and return false.
//
// TODO: Send 415 if Content-Type does not match JSON.
func (c *Context) Parse(v interface{}) bool {
	var err error

	switch c.request.Header.Get("Content-Type") {
	// lets just assume json in every case...
	//case "application/json":
	default:
		err = json.NewDecoder(c.request.Body).Decode(v)
	}

	if err != nil {
		c.Error(err, http.StatusBadRequest)

		return false
	}

	return true
}

// Responds a value by encoding it as json.
func (c *Context) Respond(v interface{}, s int) bool {
	var err error

	c.writer.WriteHeader(s)

	if v != nil {
		err = json.NewEncoder(c.writer).Encode(v)
	}

	return err == nil
}
