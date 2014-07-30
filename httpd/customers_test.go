package httpd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

func TestCustomers(t *testing.T) {
	m := &model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bob Marley",
		Email: "bob@yahoo.com",
		Address: model.Address{
			City: "Honolulu",
		},
	}

	check.Suite(&CustomersSuite{
		Url: "localhost/test",
		Tests: []CustomersTest{
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
				Status: http.StatusBadRequest,
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
		Model: m,
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
	Url        string
	Model      *model.Customer
	Store      *store.Store
	Database   *mgo.Database
	Collection *mgo.Collection
	Tests      []CustomersTest
}

func (suite *CustomersSuite) SetUpSuite(c *check.C) {
	s, err := mgo.Dial(suite.Url)
	c.Assert(err, check.IsNil)

	suite.Store = store.NewStore()
	c.Assert(suite.Store.Dial(suite.Url), check.IsNil)

	suite.Database = s.DB("")
	suite.Collection = s.DB("").C("customers")
}

func (suite *CustomersSuite) SetUpTest(c *check.C) {
	c.Assert(suite.Collection.Insert(suite.Model), check.IsNil)
}

func (s *CustomersSuite) TestServeHTTP(c *check.C) {
	h := NewCustomerHandler(s.Store)

	for i, t := range s.Tests {
		var req *http.Request

		if len(t.Body) != 0 {
			req, _ = http.NewRequest(t.Method, t.Path, strings.NewReader(t.Body))
		} else {
			req, _ = http.NewRequest(t.Method, t.Path, nil)
		}

		res := httptest.NewRecorder()

		h.ServeHTTP(res, req)

		if v := res.Code; v != t.Status {
			b := res.Body.String()

			c.Errorf("Expected #%d %s %s to respond with %d but it had %d %s", i, t.Method, t.Path, t.Status, v, b)
		}
	}
}

func (suite *CustomersSuite) TearDownTest(c *check.C) {
	c.Assert(suite.Collection.DropCollection(), check.IsNil)
}

func (suite *CustomersSuite) TearDownSuite(c *check.C) {
	c.Assert(suite.Store.Close(), check.IsNil)

	suite.Database.Session.Close()
}
