package httpd

import (
	"encoding/json"
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

	b, err := json.Marshal([]model.Customer{s.model})
	c.Check(err, check.IsNil)

	c.Check(res.Code, check.Equals, http.StatusOK)
	c.Check(res.Body.String(), check.Equals, string(b)+"\n")
}

func (s *CustomersSuite) TestListSearch(c *check.C) {
	req1, _ := http.NewRequest("GET", "/customers?search=foobar", nil)
	req2, _ := http.NewRequest("GET", "/customers?search=mar", nil)
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	s.handler.ServeHTTP(res1, req1)
	s.handler.ServeHTTP(res2, req2)

	b, err := json.Marshal([]model.Customer{s.model})
	c.Check(err, check.IsNil)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusOK)
	c.Check(res1.Body.String(), check.Equals, "[]\n")
	c.Check(res2.Body.String(), check.Equals, string(b)+"\n")
}

func (s *CustomersSuite) TestShow(c *check.C) {
	req1, _ := http.NewRequest("GET", "/customers/1234", nil)
	req2, _ := http.NewRequest("GET", "/customers/"+bson.NewObjectId().Hex(), nil)
	req3, _ := http.NewRequest("GET", "/customers/"+s.model.Id.Hex(), nil)
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()
	res3 := httptest.NewRecorder()

	s.handler.ServeHTTP(res1, req1)
	s.handler.ServeHTTP(res2, req2)
	s.handler.ServeHTTP(res3, req3)

	b, err := json.Marshal(s.model)
	c.Check(err, check.IsNil)

	c.Check(res1.Code, check.Equals, http.StatusNotFound)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusOK)
	c.Check(res3.Body.String(), check.Equals, string(b)+"\n")
}

func (s *CustomersSuite) TestCreate(c *check.C) {
	req1, _ := http.NewRequest("POST", "/customers", strings.NewReader(`{
		"name": "Edison T.",
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))
	req2, _ := http.NewRequest("POST", "/customers", strings.NewReader(`{
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	s.handler.ServeHTTP(res1, req1)
	s.handler.ServeHTTP(res2, req2)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *CustomersSuite) TestUpdate(c *check.C) {
	req1, _ := http.NewRequest("PUT", "/customers/"+s.model.Id.Hex(), strings.NewReader(`{
		"id":"`+s.model.Id.Hex()+`",
		"name": "Joe Marley",
		"email": "joe@yahoo.com",
		"address": {
			"city": "Tokio"
		}
	}`))
	req2, _ := http.NewRequest("PUT", "/customers/"+s.model.Id.Hex(), strings.NewReader(`{
		"id":"`+s.model.Id.Hex()+`",
		"name": "Joe Marley",
		"address": {
			"city": "Tokio"
		}
	}`))
	req3, _ := http.NewRequest("PUT", "/customers/1234", strings.NewReader(`{
		"id":"`+s.model.Id.Hex()+`",
		"name": "Joe Marley",
		"email": "joe@yahoo.com",
		"address": {
			"city": "Tokio"
		}
	}`))
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()
	res3 := httptest.NewRecorder()

	s.handler.ServeHTTP(res1, req1)
	s.handler.ServeHTTP(res2, req2)
	s.handler.ServeHTTP(res3, req3)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
	c.Check(res3.Code, check.Equals, http.StatusNotFound)
}

func (s *CustomersSuite) TestDestroy(c *check.C) {
	req1, _ := http.NewRequest("DELETE", "/customers/"+s.model.Id.Hex(), nil)
	req2, _ := http.NewRequest("DELETE", "/customers/1234", nil)
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	s.handler.ServeHTTP(res1, req1)
	s.handler.ServeHTTP(res2, req2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res1.Body.String(), check.Equals, "")
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
	s.handler = NewCustomers(s.store)

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
