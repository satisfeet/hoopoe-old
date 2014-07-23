package httpd

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"

	"github.com/satisfeet/hoopoe/model/validation"
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

func (s *UtilSuite) TestNotFound(c *check.C) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/bla", nil)

	NotFound(res, req)

	c.Check(res.Code, check.Equals, http.StatusNotFound)
	c.Check(res.Body.String(), check.Equals, "{\"error\":\"Not Found\"}\n")
}

func (s *UtilSuite) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
}

func (s *UtilSuite) TestErrorCode(c *check.C) {
	c.Check(ErrorCode(ErrTest), check.Equals, http.StatusInternalServerError)
	c.Check(ErrorCode(mgo.ErrNotFound), check.Equals, http.StatusNotFound)
	c.Check(ErrorCode(store.ErrInvalidQuery), check.Equals, http.StatusNotFound)
	c.Check(ErrorCode(validation.Error{}), check.Equals, http.StatusBadRequest)
}
