package httpd

import (
	"errors"
	"net/http"
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

func (s *UtilSuite) TestErrorCode(c *check.C) {
	c.Check(ErrorCode(ErrTest), check.Equals, http.StatusInternalServerError)
	c.Check(ErrorCode(store.ErrInvalidQuery), check.Equals, http.StatusBadRequest)
}
