package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	"github.com/satisfeet/hoopoe/store"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	sock = store.Product{
		Id:   bson.NewObjectId(),
		Name: "Summer Sock",
		Pricing: store.ProductPricing{
			Retail: 2.99,
		},
		Variations: store.ProductVariations{
			Sizes:  []string{"42-44"},
			Colors: []string{"black", "blue"},
		},
		Description: "Nice summer socks.",
	}
)

func TestProducts(t *testing.T) {
	s := store.NewStore()
	s.Open("localhost/test")

	Convey("Given a GET request to /products", t, func() {
		req, res := NewProductsRequestResponse("GET", "/products", "")

		Convey("ServeHTTP()", func() {
			NewProductsHandler(s).ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				products := []store.Product{}

				err := json.NewDecoder(res.Body).Decode(&products)

				So(err, ShouldBeNil)
				So(products, ShouldResemble, []store.Product{sock})
			})
		})
	})
	Convey("Given a POST request to /products", t, func() {
		req, res := NewProductsRequestResponse("POST", "/products", `{
			"name": "Business Sock",
			"pricing": {
				"retail": 2.99
			},
			"variations": {
				"sizes": ["36-38", "40-42"],
				"colors": ["red"]
			},
			"description": "Red Business socks - Screaam!!"
		}`)
		req.Header.Add("Content-Type", "application/json")

		Convey("ServeHTTP()", func() {
			NewProductsHandler(s).ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				product := store.Product{}

				err := json.NewDecoder(res.Body).Decode(&product)

				So(err, ShouldBeNil)
				So(product.Id.Valid(), ShouldBeTrue)
				So(product.Name, ShouldEqual, "Business Sock")
				So(product.Pricing.Retail, ShouldEqual, 2.99)
				So(product.Variations.Sizes, ShouldResemble, []string{"36-38", "40-42"})
				So(product.Variations.Colors, ShouldResemble, []string{"red"})
				So(product.Description, ShouldEqual, "Red Business socks - Screaam!!")
			})
		})
	})
	Convey("Given a GET request to /products/<product>", t, func() {
		req, res := NewProductsRequestResponse("GET", "/products/"+sock.Id.Hex(), "")

		Convey("ServeHTTP()", func() {
			NewProductsHandler(s).ServeHTTP(res, req)

			Convey("Should respond OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should respond json", func() {
				product := store.Product{}

				err := json.NewDecoder(res.Body).Decode(&product)

				So(err, ShouldBeNil)
				So(product, ShouldResemble, sock)
			})
		})
	})

	Convey("Given a PUT request to /products/<product>", t, func() {
		sock.Pricing.Retail += 1.00
		body, _ := json.Marshal(sock)

		req, res := NewProductsRequestResponse("PUT", "/products/"+sock.Id.Hex(), string(body))
		req.Header.Add("Content-Type", "application/json")

		Convey("ServeHTTP()", func() {
			NewProductsHandler(s).ServeHTTP(res, req)

			Convey("Should respond No Content", func() {
				So(res.Code, ShouldEqual, http.StatusNoContent)
			})
			Convey("Should respond empty body", func() {
				So(res.Body.String(), ShouldEqual, "")
			})
		})
	})

	Convey("Given a DELETE request to /products/<product>", t, func() {
		req, res := NewProductsRequestResponse("DELETE", "/products/"+sock.Id.Hex(), "")

		Convey("ServeHTTP()", func() {
			NewProductsHandler(s).ServeHTTP(res, req)

			Convey("Should respond No Content", func() {
				So(res.Code, ShouldEqual, http.StatusNoContent)
			})
			Convey("Should respond empty body", func() {
				So(res.Body.String(), ShouldEqual, "")
			})
		})
	})
}

func NewProductsRequestResponse(m string, p string, b string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(m, p, bytes.NewBufferString(b))

	session, _ := mgo.Dial("localhost/test")
	session.DB("").C("products").DropCollection()
	session.DB("").C("products").Insert(sock)
	session.Close()

	return req, httptest.NewRecorder()
}
