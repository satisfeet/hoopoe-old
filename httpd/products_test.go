package httpd

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var product = model.Product{
	Id:   bson.NewObjectId(),
	Name: "Summer Socks",
	Pricing: model.Pricing{
		Retail: 299,
	},
	Variations: []model.Variation{
		model.Variation{
			Size:  "42-44",
			Color: "black",
		},
	},
	Description: "These Summer Socks will make you really really happy.",
}

type productTest struct {
	Path   string
	Method string
	Status int
	Body   string
}

var productTests = []productTest{
	productTest{
		Path:   "/products",
		Method: "GET",
		Status: http.StatusOK,
	},
	productTest{
		Path:   "/products",
		Method: "POST",
		Status: http.StatusOK,
		Body: `{
			"name": "Winter Socks",
			"pricing": {
				"retail": 499
			},
			"variations": [
				{
					"size": "42-44",
					"color": "blue"
				}
			],
			"description": "Getting cold again. Not with these socks anymore."
		}`,
	},
	productTest{
		Path:   "/products",
		Method: "POST",
		Status: http.StatusBadRequest,
		Body: `{
			"name": "No Socks",
			"pricing": {
				"retail": -299
			}
		}`,
	},
	productTest{
		Path:   "/products/" + product.Id.Hex(),
		Method: "GET",
		Status: http.StatusOK,
	},
	productTest{
		Path:   "/products/" + bson.NewObjectId().Hex(),
		Method: "GET",
		Status: http.StatusNotFound,
	},
	productTest{
		Path:   "/products/1234",
		Method: "GET",
		Status: http.StatusBadRequest,
	},
	productTest{
		Path:   "/products/" + product.Id.Hex(),
		Method: "PUT",
		Status: http.StatusNoContent,
		Body: `{
			"id": "` + product.Id.Hex() + `",
			"name": "Winter Socks",
			"pricing": {
				"retail": 599
			},
			"variations": [
				{
					"size": "44-46",
					"color": "blue"
				}
			],
			"description": "Getting cold again. Not with these socks anymore."
		}`,
	},
	productTest{
		Path:   "/products/" + product.Id.Hex(),
		Method: "PUT",
		Status: http.StatusBadRequest,
		Body: `{
			"id": "` + product.Id.Hex() + `",
			"name": "Winter Socks",
			"pricing": {
				"retail": "599"
			},
			"description": "Getting cold again. Not with these socks anymore."
		}`,
	},
	productTest{
		Path:   "/products/" + product.Id.Hex(),
		Method: "DELETE",
		Status: http.StatusNoContent,
	},
	productTest{
		Path:   "/products/" + bson.NewObjectId().Hex(),
		Method: "DELETE",
		Status: http.StatusNotFound,
	},
}

func (s *Suite) TestProductHandler(c *check.C) {
	coll := s.db.C("products")

	for i, t := range productTests {
		c.Assert(coll.Insert(product), check.IsNil)

		var req *http.Request

		if len(t.Body) != 0 {
			req, _ = http.NewRequest(t.Method, t.Path, strings.NewReader(t.Body))
		} else {
			req, _ = http.NewRequest(t.Method, t.Path, nil)
		}

		res := httptest.NewRecorder()

		NewProductHandler(s.db).ServeHTTP(res, req)

		if v := res.Code; v != t.Status {
			b := res.Body.String()

			c.Errorf("Expected #%d %s %s to respond with %d but it had %d %s", i, t.Method, t.Path, t.Status, v, b)
		}

		_, err := coll.RemoveAll(nil)
		c.Assert(err, check.IsNil)
	}
}
