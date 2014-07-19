package context

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	Error = errors.New("test error")

	ErrorBody1 = "{\"error\":\"Not Found\"}\n"
	ErrorBody2 = "{\"error\":\"test error\"}\n"
)

func TestContextSet(t *testing.T) {
	res := httptest.NewRecorder()

	ctx := &Context{
		writer: res,
	}
	ctx.Set("X-Response-Time", "10ms")

	if h := res.Header().Get("X-Response-Time"); h != "10ms" {
		t.Errorf("Expected response header to be set but it was %s.\n", h)
	}
}

func TestContextGet(t *testing.T) {
	req, _ := http.NewRequest("POST", "/?foo=bar", nil)
	req.Header.Set("Content-Type", "application/json")

	ctx := &Context{
		request: req,
	}

	if h := ctx.Get("Content-Type"); h != "application/json" {
		t.Errorf("Expected request header to be application/json but it was %s.\n", h)
	}
}

func TestContextQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?foo=bar", nil)

	ctx := &Context{
		request: req,
	}

	if v := ctx.Query("foo"); v != "bar" {
		t.Errorf("Expected bar but had %s\n", v)
	}
}

func TestContextParse(t *testing.T) {
	m := make(map[string]string)

	req, _ := http.NewRequest("POST", "/", strings.NewReader(`
		{"foo":"bar"}
	`))
	req.Header.Add("Content-Type", "application/json")

	ctx := &Context{
		request: req,
	}

	if !ctx.Parse(&m) {
		t.Errorf("Expected to return true.\n")
	}
	if v := m["foo"]; v != "bar" {
		t.Errorf("Expected map foo to be bar but it was %s.\n", v)
	}
}

func TestContextError(t *testing.T) {
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	ctx1 := &Context{
		writer: res1,
	}
	ctx2 := &Context{
		writer: res2,
	}

	ctx1.Error(nil, http.StatusNotFound)
	ctx2.Error(Error, http.StatusBadRequest)

	if v := res1.Code; v != http.StatusNotFound {
		t.Errorf("Expected response status 404 but had %d\n", v)
	}
	if v := res2.Code; v != http.StatusBadRequest {
		t.Errorf("Expected response status 400 but had %d\n", v)
	}
	if v := res1.Body.String(); v != ErrorBody1 {
		t.Errorf("Expected response body %s but had %s\n", ErrorBody1, v)
	}
	if v := res2.Body.String(); v != ErrorBody2 {
		t.Errorf("Expected response body %s but had %s\n", ErrorBody2, v)
	}
}

func TestContextRespond(t *testing.T) {
	res := httptest.NewRecorder()

	ctx := &Context{
		writer: res,
	}
	ctx.Respond(map[string]string{"foo": "bar"}, http.StatusBadRequest)

	if v := res.Code; v != http.StatusBadRequest {
		t.Errorf("Expected response status to be 400 but it was %d.\n", v)
	}
	if v := res.Body.String(); !strings.Contains(v, `{"foo":"bar"}`) {
		t.Errorf("Expected response body to contain map but it had %s.\n", v)
	}
}
