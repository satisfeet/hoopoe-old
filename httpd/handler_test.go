package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	AuthArguments = []map[string]string{
		map[string]string{},
		map[string]string{"header": "Basic"},
		map[string]string{"username": "foo", "password": "foobar"},
	}
)

func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})
}

func TestAuth(t *testing.T) {
	req, res := NewAuthRequestResponse(Username, Password, "")

	Auth(Hello()).ServeHTTP(res, req)

	if v := res.Code; v != http.StatusOK {
		t.Errorf("Expected status to be 200 but it was %d\n", v)
	}
	if v := res.Body.String(); v != "Hello World" {
		t.Errorf("Expected body to contain hello world but it had %s\n", v)
	}

	for _, v := range AuthArguments {
		req, res := NewAuthRequestResponse(v["user"], v["pass"], v["header"])

		Auth(Hello()).ServeHTTP(res, req)

		if v := res.Code; v != http.StatusUnauthorized {
			t.Errorf("Expected status to be 401 but it was %d\n", v)
		}
		if v := res.Body.String(); v != "{\"error\":\"Unauthorized\"}\n" {
			t.Errorf("Expected body to contain unauthorized but it had %s\n", v)
		}
	}
}

func TestNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	NotFound().ServeHTTP(res, req)

	if v := res.Code; v != http.StatusNotFound {
		t.Errorf("Expected status to be 404 but it was %d.\n", v)
	}
	if v := res.Body.String(); v != "{\"error\":\"Not Found\"}\n" {
		t.Errorf("Expected body to contain error but it had %s.\n", v)
	}
}

func NewAuthRequestResponse(u, p, h string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "/", nil)

	if len(h) != 0 {
		req.Header.Set("Authorization", h)
	}
	if len(u) != 0 && len(p) != 0 {
		req.SetBasicAuth(u, p)
	}

	return req, httptest.NewRecorder()
}
