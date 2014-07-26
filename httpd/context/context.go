package context

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Context mediates to a http Request response pair but adds shortcuts and
// convenient handler for use in REST APIs.
type Context struct {
	Params   map[string]string
	Request  *http.Request
	Response http.ResponseWriter
}

// Sets the header of a response.
func (c *Context) Set(k, v string) {
	c.Response.Header().Set(k, v)
}

// Returns the header of a Request.
func (c *Context) Get(k string) string {
	return c.Request.Header.Get(k)
}

// Returns a query string value by key.
func (c *Context) Query(k string) string {
	return c.Request.URL.Query().Get(k)
}

// Returns route parameter
func (c *Context) Param(k string) string {
	return c.Params[k]
}

// Responds an json encoded error. If no error is provided will use the standard
// http status text for the given error code.
func (c *Context) Error(err error, s int) {
	if err == nil {
		err = errors.New(http.StatusText(s))
	}

	// TODO: check if Request accepts json
	http.Error(c.Response, `{"error":"`+err.Error()+`"}`, s)
}

// Parses a Request body and maps the data to the provided value. If an error
// occurs it will respond an error and return false.
//
// TODO: Return error on invalid content type.
func (c *Context) Parse(v interface{}) error {
	switch c.Request.Header.Get("Content-Type") {
	// lets just assume json in every case...
	//case "application/json":
	default:
		return json.NewDecoder(c.Request.Body).Decode(v)
	}
	return nil
}

// Responds a value by encoding it as json.
func (c *Context) Respond(v interface{}, s int) bool {
	var err error

	c.Response.WriteHeader(s)

	if v != nil {
		err = json.NewEncoder(c.Response).Encode(v)
	}

	return err == nil
}
