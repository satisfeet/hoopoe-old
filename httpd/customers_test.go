package httpd

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/satisfeet/hoopoe/store"
	. "github.com/smartystreets/goconvey/convey"
)

type Suite struct {
	Ident     string
	Fixture   bson.M
	Session   *mgo.Session
	Customers *Customers
}

func (s *Suite) SetUp() {
	store.Open(map[string]string{
		"mongo": "localhost/test",
	})

	s.Fixture = bson.M{
		"name":  "Bodo Kaiser",
		"email": "i@bodokaiser.io",
		"address": bson.M{
			"city": "Berlin",
		},
	}

	s.Session, _ = mgo.Dial("localhost/test")
	s.Session.DB("").C("customers").Insert(&s.Fixture)
	s.Session.DB("").C("customers").Find(nil).One(&s.Fixture)

	s.Ident = s.Fixture["_id"].(bson.ObjectId).Hex()

	s.Customers = &Customers{}
}

func (s *Suite) Request(m string, p string, b io.Reader) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()
	req, err := http.NewRequest(m, p, b)

	So(err, ShouldBeNil)

	s.Customers.ServeHTTP(rec, req)

	return rec
}

func TestCustomersList(t *testing.T) {
	s := &Suite{}

	Convey("GET /customers", t, func() {
		s.SetUp()

		rec := s.Request("GET", "/customers", nil)

		Convey("Should respond", func() {
			Convey("Status OK", func() {
				So(rec.Code, ShouldEqual, 200)
			})
			Convey("Body json", func() {
				var j []map[string]interface{}

				err := json.NewDecoder(rec.Body).Decode(&j)

				So(err, ShouldBeNil)
				So(j[0]["id"], ShouldEqual, s.Ident)
				So(j[0]["name"], ShouldEqual, s.Fixture["name"])
				So(j[0]["email"], ShouldEqual, s.Fixture["email"])

				a := j[0]["address"].(map[string]interface{})
				b := s.Fixture["address"].(bson.M)

				So(a["city"], ShouldResemble, b["city"])
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersShow(t *testing.T) {
	s := &Suite{}

	Convey("GET /customers/:customer", t, func() {
		s.SetUp()

		rec := s.Request("GET", "/customers/"+s.Ident, nil)

		Convey("Should respond", func() {
			Convey("Status: OK", func() {
				So(rec.Code, ShouldEqual, 200)
			})
			Convey("Body: json", func() {
				var j map[string]interface{}

				err := json.NewDecoder(rec.Body).Decode(&j)

				So(err, ShouldBeNil)
				So(j["id"], ShouldEqual, s.Ident)
				So(j["name"], ShouldEqual, s.Fixture["name"])
				So(j["email"], ShouldEqual, s.Fixture["email"])

				a := j["address"].(map[string]interface{})
				b := s.Fixture["address"].(bson.M)

				So(a["city"], ShouldResemble, b["city"])
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersCreate(t *testing.T) {
	s := &Suite{}

	Convey("POST /customers", t, func() {
		s.SetUp()

		rec := s.Request("POST", "/customers", bytes.NewBufferString(`
		{
			"name": "Haci Erdal",
			"email": "haci@erdal.de",
			"address": {
				"city": "Berlin"
			}
		}
		`))

		Convey("Should respond", func() {
			Convey("Status: OK", func() {
				So(rec.Code, ShouldEqual, 200)
			})
			Convey("Body: json", func() {
				var j map[string]interface{}

				err := json.NewDecoder(rec.Body).Decode(&j)

				So(err, ShouldBeNil)
				So(j["id"], ShouldNotBeEmpty)
				So(j["name"], ShouldEqual, "Haci Erdal")
				So(j["email"], ShouldEqual, "haci@erdal.de")

				a := j["address"].(map[string]interface{})

				So(a["city"], ShouldResemble, "Berlin")
			})
		})
		Convey("Should create customer", func() {
			res, err := s.Session.DB("").C("customers").Find(bson.M{
				"name": "Haci Erdal",
			}).Count()

			So(err, ShouldBeNil)
			So(res, ShouldEqual, 1)
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersUpdate(t *testing.T) {
	s := &Suite{}

	Convey("PUT /customers/:customer", t, func() {
		s.SetUp()

		rec := s.Request("PUT", "/customers/"+s.Ident, bytes.NewBufferString(`
		{
			"name": "Bodo Kaiser",
			"email": "bodo.kaiser@me.com",
			"address": {
				"city": "München"
			}
		}
		`))

		Convey("Should respond", func() {
			Convey("Status: No Content", func() {
				So(rec.Code, ShouldEqual, 204)
			})
		})
		Convey("Should create customer", func() {
			var j map[string]interface{}

			err := s.Session.DB("").C("customers").FindId(bson.ObjectIdHex(s.Ident)).One(&j)

			So(err, ShouldBeNil)
			So(j["name"], ShouldEqual, "Bodo Kaiser")
			So(j["email"], ShouldEqual, "bodo.kaiser@me.com")

			a := j["address"].(map[string]interface{})

			So(a["city"], ShouldResemble, "München")
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersDestroy(t *testing.T) {
	s := &Suite{}

	Convey("DELETE /customers/:customer", t, func() {
		s.SetUp()

		rec := s.Request("DELETE", "/customers/"+s.Ident, nil)

		Convey("Should respond", func() {
			Convey("Status: No Content", func() {
				So(rec.Code, ShouldEqual, 204)
			})
		})
		Convey("Should remove customer", func() {
			res, err := s.Session.DB("").C("customers").FindId(
				bson.ObjectIdHex(s.Ident),
			).Count()

			So(err, ShouldBeNil)
			So(res, ShouldEqual, 0)
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func (s *Suite) TearDown() {
	s.Session.DB("").C("customers").RemoveAll(bson.M{})
	s.Session.Close()
}
