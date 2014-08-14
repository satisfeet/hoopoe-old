package httpd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"

	"github.com/satisfeet/go-context"
)

var hs = &HandlerSuite{
	url: "localhost/test",
}

func TestHandler(t *testing.T) {
	check.Suite(hs)
	check.TestingT(t)
}

type HandlerSuite struct {
	url      string
	session  *mgo.Session
	database *mgo.Database
}

func (s *HandlerSuite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.session = sess
	s.database = sess.DB("")
}

func (s *HandlerSuite) TearDownSuite(c *check.C) {
	s.session.Close()
}

func ctx(m, p string, b io.Reader) (*context.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(m, p, b)

	return &context.Context{
		Response: res,
		Request:  req,
	}, res
}
