package httpd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/go-context"
	"github.com/satisfeet/hoopoe/store/mongo"
)

func TestHttpd(t *testing.T) {
	check.Suite(&Suite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type Suite struct {
	url   string
	mongo *mongo.Store
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
	s.mongo = &mongo.Store{}
	err := s.mongo.Dial(s.url)
	c.Assert(err, check.IsNil)
}

func (s *Suite) SetUpTest(c *check.C) {
	c.Assert(s.mongo.Insert("products", &product), check.IsNil)
	c.Assert(s.mongo.Insert("customers", &customer), check.IsNil)

	f, err := s.mongo.CreateFile("products")
	c.Assert(err, check.IsNil)
	product.Images = []bson.ObjectId{f.Id}

	_, err = f.Write([]byte("Hello World"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
}

func (s *Suite) TearDownSuite(c *check.C) {
	c.Assert(s.mongo.Close(), check.IsNil)
}

func (s *Suite) TearDownTest(c *check.C) {
	c.Assert(s.mongo.RemoveAll("products", nil), check.IsNil)
	c.Assert(s.mongo.RemoveAll("products.files", nil), check.IsNil)
	c.Assert(s.mongo.RemoveAll("products.chunks", nil), check.IsNil)
	c.Assert(s.mongo.RemoveAll("customers", nil), check.IsNil)
}
