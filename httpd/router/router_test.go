package router

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/satisfeet/hoopoe/httpd/context"
)

func TestRouter(t *testing.T) {
	req1, res1 := NewRequestResponse("GET", "/echo/bodo")
	req2, res2 := NewRequestResponse("PUT", "/hello")

	r := NewRouter()
	r.Handle("GET", "/echo/:name", &EchoHandler{})
	r.HandleFunc("PUT", "/hello", HelloHandlerFunc)

	r.ServeHTTP(res1, req1)
	if v := res1.Code; v != http.StatusOK {
		t.Errorf("Expected status OK but had %d.\n", v)
	}
	if v := res1.Body.String(); !strings.Contains(v, `{"message":"bodo"}`) {
		t.Errorf("Expected status message body but had %s.\n", v)
	}

	r.ServeHTTP(res2, req2)
	if v := res2.Code; v != http.StatusOK {
		t.Errorf("Expected status OK but had %d.\n", v)
	}
	if v := res2.Body.String(); !strings.Contains(v, `{"message":"hello"}`) {
		t.Errorf("Expected status message body but had %s.\n", v)
	}
}

type EchoHandler struct{}

func (h *EchoHandler) ServeHTTP(c *context.Context) {
	c.Respond(map[string]string{
		"message": c.Param("name"),
	}, http.StatusOK)
}

func HelloHandlerFunc(c *context.Context) {
	c.Respond(map[string]string{
		"message": "hello",
	}, http.StatusOK)
}

func NewRequestResponse(m, p string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(m, p, nil)

	return req, httptest.NewRecorder()
}
