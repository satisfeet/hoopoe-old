package httpd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type customerTest struct {
	Status  int
	Method  string
	Path    string
	Body    string
	ResBody string
}

var (
	customer = model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bob Marley",
		Email: "bob@yahoo.com",
		Address: model.Address{
			City: "Honolulu",
		},
	}

	customerTests = []customerTest{
		customerTest{
			Status: http.StatusOK,
			Method: "GET",
			Path:   "/customers",
		},
		customerTest{
			Status: http.StatusOK,
			Method: "GET",
			Path:   "/customers?search=Berl",
		},
		customerTest{
			Status: http.StatusOK,
			Method: "POST",
			Path:   "/customers",
			Body: `{
				"name":"Edison Trent",
				"email":"edison@apache.org",
				"address":{
					"city":"Leeds"
				}
			}`,
		},
		customerTest{
			Status: http.StatusBadRequest,
			Method: "GET",
			Path:   "/customers/1234",
		},
		customerTest{
			Status: http.StatusNotFound,
			Method: "GET",
			Path:   "/customers/" + bson.NewObjectId().Hex(),
		},
		customerTest{
			Status: http.StatusOK,
			Method: "GET",
			Path:   "/customers/" + customer.Id.Hex(),
		},
		customerTest{
			Status: http.StatusNoContent,
			Method: "PUT",
			Path:   "/customers/" + customer.Id.Hex(),
			Body: `{
				"id":"` + customer.Id.Hex() + `",
				"name":"Bob Marley",
				"email":"bob@yahoo.com",
				"address":{
					"city":"Honolulu"
				}
			}`,
		},
		customerTest{
			Status: http.StatusNoContent,
			Method: "DELETE",
			Path:   "/customers/" + customer.Id.Hex(),
		},
	}
)

func TestCustomers(t *testing.T) {
	store.Open("localhost/test")
	defer store.Close()

	c := &Customers{
		Store: &store.Store{
			Name: "customers",
		},
	}

	s, _ := mgo.Dial("localhost/test")
	defer s.Close()

	for _, ct := range customerTests {
		s.DB("").C("customers").Insert(&customer)

		req, _ := http.NewRequest(ct.Method, ct.Path, strings.NewReader(ct.Body))
		res := httptest.NewRecorder()

		c.ServeHTTP(res, req)

		if v := res.Code; v != ct.Status {
			t.Errorf("Expected %s %s status to be %d but it was %d with %s.\n", ct.Method, ct.Path, ct.Status, v, res.Body.String())
		}

		s.DB("").C("customers").DropCollection()
	}
}
