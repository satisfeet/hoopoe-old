package httpd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/store/mongo"
)

var hs = &HandlerSuite{
	url: "localhost/test",
}

func TestHandler(t *testing.T) {
	check.Suite(hs)
	check.TestingT(t)
}

type HandlerSuite struct {
	url   string
	mongo *mongo.Store
}

func (s *HandlerSuite) SetUpSuite(c *check.C) {
	s.mongo = &mongo.Store{}

	err := s.mongo.Dial(s.url)
	c.Assert(err, check.IsNil)
}

func (s *HandlerSuite) TearDownSuite(c *check.C) {
	c.Assert(s.mongo.Close(), check.IsNil)
}

func ctx(m, p string, b io.Reader) (*context.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(m, p, b)

	return &context.Context{
		Response: res,
		Request:  req,
	}, res
}
