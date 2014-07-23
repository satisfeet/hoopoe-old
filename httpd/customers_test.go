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
	m := model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bob Marley",
		Email: "bob@yahoo.com",
		Address: model.Address{
			City: "Honolulu",
		},
	}

	check.Suite(&CustomersSuite{
		url: "localhost/test",
		tests: []CustomersTest{
			CustomersTest{
				Path:   "/customers",
				Method: "GET",
				Status: http.StatusOK,
			},
			CustomersTest{
				Path:   "/customers?search=mar",
				Method: "GET",
				Status: http.StatusOK,
			},
			CustomersTest{
				Path:   "/customers?search=foobar",
				Method: "GET",
				Status: http.StatusOK,
			},
			CustomersTest{
				Path:   "/customers",
				Method: "POST",
				Status: http.StatusOK,
				Body: `{
					"name": "Edison T.",
					"email": "edison@t.com",
					"address": {
						"city": "Leeds"
					}
				}`,
			},
			CustomersTest{
				Path:   "/customers",
				Method: "POST",
				Status: http.StatusBadRequest,
				Body: `{
					"email": "edison@t.com",
					"address": {
						"city": "Leeds"
					}
				}`,
			},
			CustomersTest{
				Path:   "/customers/" + m.Id.Hex(),
				Method: "GET",
				Status: http.StatusOK,
			},
			CustomersTest{
				Path:   "/customers/" + bson.NewObjectId().Hex(),
				Method: "GET",
				Status: http.StatusNotFound,
			},
			CustomersTest{
				Path:   "/customers/1234",
				Method: "GET",
				Status: http.StatusNotFound,
			},
			CustomersTest{
				Path:   "/customers/" + m.Id.Hex(),
				Method: "PUT",
				Status: http.StatusNoContent,
				Body: `{
					"id": "` + m.Id.Hex() + `",
					"name": "Bob Marley",
					"email": "bob@marley.com",
					"address": {
						"city": "New York"
					}
				}`,
			},
			CustomersTest{
				Path:   "/customers/" + m.Id.Hex(),
				Method: "PUT",
				Status: http.StatusBadRequest,
				Body: `{
					"id": "` + m.Id.Hex() + `",
					"name": "Bob Marley",
					"address": {
						"city": "New York"
					}
				}`,
			},
			CustomersTest{
				Path:   "/customers/" + m.Id.Hex(),
				Method: "DELETE",
				Status: http.StatusNoContent,
			},
			CustomersTest{
				Path:   "/customers/" + bson.NewObjectId().Hex(),
				Method: "DELETE",
				Status: http.StatusNotFound,
			},
		},
		model: m,
	})
	check.TestingT(t)
}

type CustomersTest struct {
	Path   string
	Method string
	Status int
	Body   string
}

type CustomersSuite struct {
	url     string
	model   model.Customer
	store   *store.Store
	session *store.Session
	handler *Customers
	tests   []CustomersTest
}

func (s *CustomersSuite) TestServeHTTP(c *check.C) {
	for i, t := range s.tests {
		var req *http.Request

		if len(t.Body) != 0 {
			req, _ = http.NewRequest(t.Method, t.Path, strings.NewReader(t.Body))
		} else {
			req, _ = http.NewRequest(t.Method, t.Path, nil)
		}

		res := httptest.NewRecorder()

		s.handler.ServeHTTP(res, req)

		if v := res.Code; v != t.Status {
			b := res.Body.String()

			c.Errorf("Expected #%d %s %s to respond with %d but it had %d %s", i, t.Method, t.Path, t.Status, v, b)
		}
	}
}

func (s *CustomersSuite) SetUpSuite(c *check.C) {
	s.store = &store.Store{
		Name:    "customers",
		Session: &store.Session{},
	}
	s.handler = NewCustomers(s.store)

	c.Assert(s.store.Session.Open(s.url), check.IsNil)
}

func (s *CustomersSuite) SetUpTest(c *check.C) {
	m := s.store.Session.Mongo()
	defer m.Close()

	c.Assert(m.DB("").C("customers").Insert(&s.model), check.IsNil)
}

func (s *CustomersSuite) TearDownTest(c *check.C) {
	m := s.store.Session.Mongo()
	defer m.Close()

	c.Assert(m.DB("").C("customers").DropCollection(), check.IsNil)
}

func (s *CustomersSuite) TearDownSuite(c *check.C) {
	s.store.Session.Close()
}
