package context

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	Error = errors.New("test error")

	ErrorBody1 = "{\"error\":\"Not Found\"}\n"
	ErrorBody2 = "{\"error\":\"test error\"}\n"
)

func TestContextQuery(t *testing.T) {
	req, _ := http.NewRequest("GET", "/?foo=bar", nil)

	ctx := &Context{
		request: req,
	}

	if v := ctx.Query("foo"); v != "bar" {
		t.Errorf("Expected bar but had %s\n", v)
	}
}

func TestError(t *testing.T) {
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
