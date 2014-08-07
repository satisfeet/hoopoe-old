package httpd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satisfeet/go-context"
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
)

func TestHttpd(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	url string
	db  *mgo.Database
}

func (s *Suite) Context(m, p string, b io.Reader) (*context.Context, *httptest.ResponseRecorder) {
	res := httptest.NewRecorder()
	req, _ := http.NewRequest(m, p, b)

	return &context.Context{
		Response: res,
		Request:  req,
	}, res
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.db = sess.DB("")
}

func (s *Suite) SetUpTest(c *check.C) {
	c.Assert(s.db.C("products").Insert(product), check.IsNil)
	c.Assert(s.db.C("customers").Insert(customer), check.IsNil)
}

func (s *Suite) TearDownSuite(c *check.C) {
	s.db.Session.Close()
}

func (s *Suite) TearDownTest(c *check.C) {
	var err error

	_, err = s.db.C("products").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.db.C("customers").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}
