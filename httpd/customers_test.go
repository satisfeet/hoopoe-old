package httpd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

func TestCustomers(t *testing.T) {
	check.Suite(&CustomersSuite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type CustomersSuite struct {
	url     string
	model   model.Customer
	store   *store.Store
	session *store.Session
	handler *Customers
}

func (s *CustomersSuite) TestList(c *check.C) {
	req, _ := http.NewRequest("GET", "/customers", nil)
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusOK)
}

func (s *CustomersSuite) TestShow(c *check.C) {
	req, _ := http.NewRequest("GET", "/customers/"+s.model.Id.Hex(), nil)
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusOK)
}

func (s *CustomersSuite) TestCreate(c *check.C) {
	req, _ := http.NewRequest("POST", "/customers", strings.NewReader(`{
		"name": "Edison T.",
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusOK)
}

func (s *CustomersSuite) TestUpdate(c *check.C) {
	req, _ := http.NewRequest("PUT", "/customers/"+s.model.Id.Hex(), strings.NewReader(`{
		"id":"`+s.model.Id.Hex()+`",
		"name": "Joe Marley",
		"email": "joe@yahoo.com",
		"address": {
			"city": "Tokio"
		}
	}`))
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusNoContent)
}

func (s *CustomersSuite) TestDestroy(c *check.C) {
	req, _ := http.NewRequest("DELETE", "/customers/"+s.model.Id.Hex(), nil)
	res := httptest.NewRecorder()

	s.handler.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusNoContent)

}

func (s *CustomersSuite) SetUpSuite(c *check.C) {
	s.model = model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bob Marley",
		Email: "bob@yahoo.com",
		Address: model.Address{
			City: "Honolulu",
		},
	}
	s.session = &store.Session{}
	s.store = &store.Store{
		Name:    "customers",
		Session: s.session,
	}
	s.handler = &Customers{
		Store: s.store,
	}

	c.Assert(s.session.Open(s.url), check.IsNil)
}

func (s *CustomersSuite) SetUpTest(c *check.C) {
	m := s.session.Mongo()
	defer m.Close()

	c.Assert(m.DB("").C("customers").Insert(&s.model), check.IsNil)
}

func (s *CustomersSuite) TearDownTest(c *check.C) {
	m := s.session.Mongo()
	defer m.Close()

	c.Assert(m.DB("").C("customers").DropCollection(), check.IsNil)
}

func (s *CustomersSuite) TearDownSuite(c *check.C) {
	s.session.Close()
}
