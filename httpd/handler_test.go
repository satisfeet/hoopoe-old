package httpd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})
}

func TestAuth(t *testing.T) {
	req, res := NewAuthRequestResponse()
	req.SetBasicAuth(Username, Password)

	Auth(Hello()).ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Error("Expected status to be 200 but it was %d\n", res.Code)
	}
	if res.Body.String() != "Hello World" {
		t.Error("Expected body to contain hello world but it had %s\n", res.Body.String())
	}
}

func TestAuthWithoutHeader(t *testing.T) {
	req, res := NewAuthRequestResponse()

	Auth(Hello()).ServeHTTP(res, req)

	if res.Code != http.StatusUnauthorized {
		t.Error("Expected status to be 401 but it was %d\n", res.Code)
	}
	if !strings.Contains(res.Body.String(), http.StatusText(http.StatusUnauthorized)) {
		t.Error("Expected body to contain unauthorized but it had %s\n", res.Body.String())
	}
}

func TestAuthWithInvalidHeader(t *testing.T) {
	req, res := NewAuthRequestResponse()
	req.Header.Set("Authorization", "Basic")

	Auth(Hello()).ServeHTTP(res, req)

	if res.Code != http.StatusUnauthorized {
		t.Error("Expected status to be 401 but it was %d\n", res.Code)
	}
	if !strings.Contains(res.Body.String(), http.StatusText(http.StatusUnauthorized)) {
		t.Error("Expected body to contain unauthorized but it had %s\n", res.Body.String())
	}
}

func TestAuthWithInvalidCredentials(t *testing.T) {
	req, res := NewAuthRequestResponse()
	req.SetBasicAuth("foo", "foobar")

	Auth(Hello()).ServeHTTP(res, req)

	if res.Code != http.StatusUnauthorized {
		t.Error("Expected status to be 401 but it was %d\n", res.Code)
	}
	if !strings.Contains(res.Body.String(), http.StatusText(http.StatusUnauthorized)) {
		t.Error("Expected body to contain unauthorized but it had %s\n", res.Body.String())
	}
}

func NewAuthRequestResponse() (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "/", nil)

	return req, httptest.NewRecorder()
}
