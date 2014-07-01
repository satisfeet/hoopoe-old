package httpd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/satisfeet/hoopoe/httpd/router"
)

type Suite struct {
	db     *mgo.Database
	router *router.Router
	Server *httptest.Server
}

func (s *Suite) SetUp() {
	session, err := mgo.Dial("localhost")

	So(err, ShouldBeNil)

	s.db = session.DB("test")

	err = s.db.C("customers").Insert(bson.M{
		"name":  "Bodo Kaiser",
		"email": "i@bodokaiser.io",
		"address": bson.M{
			"city": "Berlin",
		},
	})

	So(err, ShouldBeNil)

	s.router = router.New()

	CustomersInit(s.router)

	s.Server = httptest.NewServer(s.router)
}

func TestCustomersList(t *testing.T) {
	Convey("GET /customers", t, func() {
		s := &Suite{}
		s.SetUp()

		r, err := http.Get(s.Server.URL + "/customers")

		Convey("Should have no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Should have response body", func() {
			var j []map[string]interface{}

			Convey("Which encodes as json", func() {
				err := json.NewDecoder(r.Body).Decode(&j)

				Convey("Without error", func() {
					So(err, ShouldBeNil)
				})
				Convey("As array", func() {
					So(j[0]["name"], ShouldEqual, "Bodo Kaiser")
					So(j[0]["email"], ShouldEqual, "i@bodokaiser.io")

					a := j[0]["address"].(map[string]interface{})

					So(a["city"], ShouldEqual, "Berlin")
				})
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersShow(t *testing.T) {
	Convey("GET /customers/:customer", t, func() {
		s := &Suite{}
		s.SetUp()

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersCreate(t *testing.T) {
	Convey("POST /customers", t, func() {
		s := &Suite{}
		s.SetUp()

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersUpdate(t *testing.T) {
	Convey("PUT /customers/:customer", t, func() {
		s := &Suite{}
		s.SetUp()

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersDestroy(t *testing.T) {
	Convey("DELETE /customers/:customer", t, func() {
		s := &Suite{}
		s.SetUp()

		Reset(func() {
			s.TearDown()
		})
	})
}

func (s *Suite) TearDown() {
	s.Server.Close()

	_, err := s.db.C("customers").RemoveAll(bson.M{})

	So(err, ShouldBeNil)

	s.db.Session.Close()
}
