package httpd

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"

	"github.com/satisfeet/hoopoe/store"
)

var (
	ErrTest = errors.New("a test error")
)

func TestUtil(t *testing.T) {
	check.Suite(&UtilSuite{})
	check.TestingT(t)
}

type UtilSuite struct{}

func (s *UtilSuite) TestAuth(c *check.C) {
	req := make([]*http.Request, 3)
	res := make([]*httptest.ResponseRecorder, 3)

	for i := 0; i < 3; i++ {
		req[i], _ = http.NewRequest("GET", "/", nil)
		res[i] = httptest.NewRecorder()
	}

	req[0].SetBasicAuth("foo", "bar")
	req[1].SetBasicAuth("foo", "bla")
	req[2].Header.Add("Authorization", "Basic ")

	Basic = "foo:bar"
	Auth(s).ServeHTTP(res[0], req[0])
	Auth(s).ServeHTTP(res[1], req[1])
	Auth(s).ServeHTTP(res[2], req[2])

	c.Check(res[0].Code, check.Equals, http.StatusOK)
	c.Check(res[1].Code, check.Equals, http.StatusUnauthorized)
	c.Check(res[2].Code, check.Equals, http.StatusUnauthorized)

	c.Check(res[0].Body.String(), check.Equals, "Hello World")
}

func (s *UtilSuite) TestNotFound(c *check.C) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bla", nil)

	NotFound().ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusNotFound)
	c.Check(res.Body.String(), check.Equals, "{\"error\":\"Not Found\"}\n")
}

func (s *UtilSuite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func (s *UtilSuite) TestErrorCode(c *check.C) {
	c.Check(ErrorCode(ErrTest), check.Equals, http.StatusInternalServerError)
	c.Check(ErrorCode(store.ErrInvalidQuery), check.Equals, http.StatusBadRequest)
}