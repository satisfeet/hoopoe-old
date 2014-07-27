package httpd

import (
	"errors"
	"net/http"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-validation"
)

func TestHttpd(t *testing.T) {
	check.Suite(&HttpdSuite{})
	check.TestingT(t)
}

type HttpdSuite struct{}

func (s *HttpdSuite) TestErrorCode(c *check.C) {
	c.Check(ErrorCode(errors.New("test")), check.Equals, http.StatusInternalServerError)
	c.Check(ErrorCode(mgo.ErrNotFound), check.Equals, http.StatusNotFound)
	c.Check(ErrorCode(validation.Error{}), check.Equals, http.StatusBadRequest)
}
