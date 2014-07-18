package httpd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	url = "localhost/test"

	customer1 = model.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bodo Kaiser",
		Email: "i@bodokaiser.io",
		Address: model.Address{
			City:    "Berlin",
			Street:  "Geiserichstr. 3",
			Zipcode: 12105,
		},
	}
	customer2 = map[string]interface{}{
		"name":  "Edison Trent",
		"email": "edison@me.com",
		"address": map[string]string{
			"city": "Leeds",
		},
	}

	customer1JSON, _ = json.Marshal(customer1)
	customer2JSON, _ = json.Marshal(customer2)
)

func TestCustomersHandler(t *testing.T) {
	store.Open(url)
	defer store.Close()

	s, _ := mgo.Dial(url)
	_, err1 := s.DB("").C(model.CustomerName).RemoveAll(bson.M{})
	err2 := s.DB("").C(model.CustomerName).Insert(&customer1)
	defer s.Close()

	if err1 != nil {
		panic(err1)
	}
	if err2 != nil {
		panic(err2)
	}

	Convey("Given a GET request to /customers", t, func() {
		req, res := NewCustomersRequestResponse("GET", "/customers", "")

		Convey("ServeHTTP()", func() {
			NewCustomersHandler().ServeHTTP(res, req)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 200)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, string(customer1JSON))
			})
		})
	})
	Convey("Given a POST request to /customers", t, func() {
		req, res := NewCustomersRequestResponse("POST", "/customers", string(customer2JSON))

		Convey("ServeHTTP()", func() {
			NewCustomersHandler().ServeHTTP(res, req)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 200)
			})
			Reset(func() {
				s.DB("").C(model.CustomerName).Remove(bson.M{"name": "Edison Trent"})
			})
		})
	})
	Convey("Given a GET request to /customers/:id", t, func() {
		req, res := NewCustomersRequestResponse("GET", "/customers/"+customer1.Id.Hex(), "")

		Convey("ServeHTTP()", func() {
			NewCustomersHandler().ServeHTTP(res, req)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 200)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, string(customer1JSON))
			})
		})
	})
	Convey("Given a PUT request to /customers/:id", t, func() {
		req, res := NewCustomersRequestResponse("PUT", "/customers/"+customer1.Id.Hex(),
			strings.Replace(string(customer1JSON), "i@bodokaiser.io", "bodo.kaiser@me.com", -1))

		Convey("ServeHTTP()", func() {
			NewCustomersHandler().ServeHTTP(res, req)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 204)
			})
		})
	})
	Convey("Given a DELETE request to /customers/:id", t, func() {
		req, res := NewCustomersRequestResponse("DELETE", "/customers/"+customer1.Id.Hex(), "")

		Convey("ServeHTTP()", func() {
			NewCustomersHandler().ServeHTTP(res, req)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, 204)
			})
		})
	})
}

func NewCustomersRequestResponse(m, s, b string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(m, s, strings.NewReader(b))

	return req, httptest.NewRecorder()
}
