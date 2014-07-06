package context

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Context is a helper struct to handle a HTTP action.
//
// It provides a lot of convenient methods which make it very easy
// to respond something, handle errors or analyze the request.
type Context struct {
	request *http.Request
	writer  http.ResponseWriter
	params  httprouter.Params
}

// Handle is a function which takes a Context as argument.
//
// Can be used over httprouter.Handle or http.HandlerFunc.
type Handle func(*Context)

// Returns an initialized Context.
//
// As we utilize httprouter for routing we also need the parsed Params
// also this allows us to hide the ugly Params interface.
func New(w http.ResponseWriter, r *http.Request, p httprouter.Params) *Context {
	return &Context{r, w, p}
}

// Wraps a context Handle to work with httprouter.
//
// This makes it easy to adapt context style handlers with httprouter.
func HandleFunc(handle Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		c := New(w, r, p)

		handle(c)
	}
}

// Returns the parsed URL parameter.
func (c *Context) Param(key string) string {
	return c.params.ByName(key)
}

// Returns the URL value assigned to the query string key.
func (c *Context) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

// Responds an HTTP error.
//
// If error is nil then we will use the default http status text
// for the corresponding code. We support this as http status codes
// are quite expressive and cover a lot common error cases.
// Furthermore we will send back the error message in the form
// requested by the client (not supported yet).
func (c *Context) Error(err error, code int) {
	if err == nil {
		err = errors.New(http.StatusText(code))
	}

	// TODO: check if request accepts json
	http.Error(c.writer, `{"error":"`+err.Error()+`"}`, code)
}

// Parses a HTTP request.
//
// This will decode the HTTP request body by using the defined content type.
// If something goes wrong or no content type matches an error is returned.
func (c *Context) Parse(value interface{}) error {
	if c.request.Header.Get("Content-Type") == "application/json" {
		return json.NewDecoder(c.request.Body).Decode(value)
	}

	return errors.New("Unsupported encoding.")
}

// Responds to a HTTP request.
//
// This will decode the given value into an accepted form.
// If an error occurs it will be returned to be handled by the user.
func (c *Context) Respond(value interface{}, code int) error {
	c.writer.WriteHeader(code)

	if value != nil {
		return json.NewEncoder(c.writer).Encode(value)
	}

	return nil
}
