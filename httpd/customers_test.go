package httpd

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/satisfeet/hoopoe/store"
)

var (
	bob = store.Customer{
		Id:    bson.NewObjectId(),
		Name:  "Bob Jersey",
		Email: "jersaybob@redneck.co",
		Address: store.CustomerAddress{
			City: "New Jersey",
		},
	}
)

func TestCustomersAPIServeHTTP(t *testing.T) {
	Convey("Given a GET request to /customers", t, func() {
		req, res, ca := requestCA("GET", "/customers", "")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				customers := []store.Customer{}

				err := json.NewDecoder(res.Body).Decode(&customers)

				So(err, ShouldBeNil)
				So(customers, ShouldResemble, []store.Customer{bob})
			})
		})
	})
	Convey("Given a GET request to /customers?search=Berl", t, func() {
		req, res, ca := requestCA("GET", "/customers?search=Berl", "")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				customers := []store.Customer{}

				err := json.NewDecoder(res.Body).Decode(&customers)

				So(err, ShouldBeNil)
				So(customers, ShouldResemble, []store.Customer{})
			})
		})
	})
	Convey("Given a GET request to /customers?search=New", t, func() {
		req, res, ca := requestCA("GET", "/customers?search=New", "")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				customers := []store.Customer{}

				err := json.NewDecoder(res.Body).Decode(&customers)

				So(err, ShouldBeNil)
				So(customers, ShouldResemble, []store.Customer{bob})
			})
		})
	})

	Convey("Given a POST request to /customers", t, func() {
		req, res, ca := requestCA("POST", "/customers", `
			{"name":"Sandra","email":"sandra@yahoo.uk","address":{"city":"Leeds"}}
		`)
		req.Header.Add("Content-Type", "application/json")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				customer := store.Customer{}

				err := json.NewDecoder(res.Body).Decode(&customer)

				So(err, ShouldBeNil)
				So(customer.Id.Valid(), ShouldBeTrue)
				So(customer.Name, ShouldEqual, "Sandra")
				So(customer.Email, ShouldEqual, "sandra@yahoo.uk")
				So(customer.Address.City, ShouldEqual, "Leeds")
			})
		})
	})

	Convey("Given a GET request to /customers/<customer>", t, func() {
		req, res, ca := requestCA("GET", "/customers/"+bob.Id.Hex(), "")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				customer := store.Customer{}

				err := json.NewDecoder(res.Body).Decode(&customer)

				So(err, ShouldBeNil)
				So(customer, ShouldResemble, bob)
			})
		})
	})

	Convey("Given a PUT request to /customers/<customer>", t, func() {
		bob.Email = "bob@gmail.com"
		body, _ := json.Marshal(bob)

		req, res, ca := requestCA("PUT", "/customers/"+bob.Id.Hex(), string(body))
		req.Header.Add("Content-Type", "application/json")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond No Content", func() {
				So(res.Code, ShouldEqual, http.StatusNoContent)
			})
			Convey("Should respond empty body", func() {
				So(res.Body.String(), ShouldEqual, "")
			})
		})
	})

	Convey("Given a DELETE request to /customers/<customer>", t, func() {
		req, res, ca := requestCA("DELETE", "/customers/"+bob.Id.Hex(), "")

		Convey("ServeHTTP()", func() {
			ca.ServeHTTP(res, req)

			Convey("Should respond No Content", func() {
				So(res.Code, ShouldEqual, http.StatusNoContent)
			})
			Convey("Should respond empty body", func() {
				So(res.Body.String(), ShouldEqual, "")
			})
		})
	})
}

func requestCA(m string, p string, b string) (*http.Request, *httptest.ResponseRecorder, *CustomerAPI) {
	req, _ := http.NewRequest(m, p, bytes.NewBufferString(b))

	s := store.NewStore()
	s.Open("localhost/test")
	s.Mongo().C("customers").DropCollection()
	s.Mongo().C("customers").Insert(bob)

	return req, httptest.NewRecorder(), &CustomerAPI{s}
}
