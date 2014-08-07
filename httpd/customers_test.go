package httpd

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var customer = model.Customer{
	Id:    bson.NewObjectId(),
	Name:  "Bob Marley",
	Email: "bob@yahoo.com",
	Address: model.Address{
		City: "Honolulu",
	},
}

type customerTest struct {
	Path   string
	Method string
	Status int
	Body   string
}

var customerTests = []customerTest{
	customerTest{
		Path:   "/customers",
		Method: "GET",
		Status: http.StatusOK,
	},
	customerTest{
		Path:   "/customers?search=mar",
		Method: "GET",
		Status: http.StatusOK,
	},
	customerTest{
		Path:   "/customers?search=foobar",
		Method: "GET",
		Status: http.StatusOK,
	},
	customerTest{
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
	customerTest{
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
	customerTest{
		Path:   "/customers/" + customer.Id.Hex(),
		Method: "GET",
		Status: http.StatusOK,
	},
	customerTest{
		Path:   "/customers/" + bson.NewObjectId().Hex(),
		Method: "GET",
		Status: http.StatusNotFound,
	},
	customerTest{
		Path:   "/customers/1234",
		Method: "GET",
		Status: http.StatusBadRequest,
	},
	customerTest{
		Path:   "/customers/" + customer.Id.Hex(),
		Method: "PUT",
		Status: http.StatusNoContent,
		Body: `{
			"id": "` + customer.Id.Hex() + `",
			"name": "Bob Marley",
			"email": "bob@marley.com",
			"address": {
				"city": "New York"
			}
		}`,
	},
	customerTest{
		Path:   "/customers/" + customer.Id.Hex(),
		Method: "PUT",
		Status: http.StatusBadRequest,
		Body: `{
			"id": "` + customer.Id.Hex() + `",
			"name": "Bob Marley",
			"address": {
				"city": "New York"
			}
		}`,
	},
	customerTest{
		Path:   "/customers/" + customer.Id.Hex(),
		Method: "DELETE",
		Status: http.StatusNoContent,
	},
	customerTest{
		Path:   "/customers/" + bson.NewObjectId().Hex(),
		Method: "DELETE",
		Status: http.StatusNotFound,
	},
}

func (s *Suite) TestCustomerHandler(c *check.C) {
	coll := s.db.C("customers")

	for i, t := range customerTests {
		c.Assert(coll.Insert(customer), check.IsNil)

		var req *http.Request

		if len(t.Body) != 0 {
			req, _ = http.NewRequest(t.Method, t.Path, strings.NewReader(t.Body))
		} else {
			req, _ = http.NewRequest(t.Method, t.Path, nil)
		}

		res := httptest.NewRecorder()

		NewCustomerHandler(s.db).ServeHTTP(res, req)

		if v := res.Code; v != t.Status {
			b := res.Body.String()

			c.Errorf("Expected #%d %s %s to respond with %d but it had %d %s", i, t.Method, t.Path, t.Status, v, b)
		}

		_, err := coll.RemoveAll(nil)
		c.Assert(err, check.IsNil)
	}
}
