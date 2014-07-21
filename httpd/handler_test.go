package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

var (
	AuthArguments = []map[string]string{
		map[string]string{},
		map[string]string{"header": "Basic"},
		map[string]string{"username": "foo", "password": "foobar"},
	}
)

func TestHandler(t *testing.T) {
	check.Suite(&HandlerSuite{})
	check.TestingT(t)
}

type HandlerSuite struct{}

func (s *HandlerSuite) TestAuth(c *check.C) {
	req := make([]*http.Request, 3)
	res := make([]*httptest.ResponseRecorder, 3)

	for i := 0; i < 3; i++ {
		req[i], _ = http.NewRequest("GET", "/", nil)
		res[i] = httptest.NewRecorder()
	}

	req[0].SetBasicAuth(Username, Password)
	req[1].SetBasicAuth(Password, Username)
	req[2].Header.Add("Authorization", "Basic ")

	Auth(s.hello()).ServeHTTP(res[0], req[0])
	Auth(s.hello()).ServeHTTP(res[1], req[1])
	Auth(s.hello()).ServeHTTP(res[2], req[2])

	c.Check(res[0].Code, check.Equals, http.StatusOK)
	c.Check(res[1].Code, check.Equals, http.StatusUnauthorized)
	c.Check(res[2].Code, check.Equals, http.StatusUnauthorized)

	c.Check(res[0].Body.String(), check.Equals, "Hello World")
}

func (s *HandlerSuite) TestNotFound(c *check.C) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bla", nil)

	NotFound().ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusNotFound)
	c.Check(res.Body.String(), check.Equals, "{\"error\":\"Not Found\"}\n")
}

func (s *HandlerSuite) hello() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World"))
	})
}
