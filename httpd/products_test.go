package httpd

import (
	"net/http"
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

func (s *Suite) TestProductHandlerList(c *check.C) {
	h := NewProductHandler(s.db)

	ctx1, res1 := s.Context("GET", "/products", nil)

	h.List(ctx1)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(strings.HasPrefix(res1.Body.String(), "[{"), check.Equals, true)
}

func (s *Suite) TestProductHandlerShow(c *check.C) {
	h := NewProductHandler(s.db)

	ctx1, res1 := s.Context("GET", "/", nil)
	ctx1.Params = map[string]string{"pid": product.Id.Hex()}
	ctx2, res2 := s.Context("GET", "/", nil)
	ctx2.Params = map[string]string{"pid": bson.NewObjectId().Hex()}
	ctx3, res3 := s.Context("GET", "/", nil)
	ctx3.Params = map[string]string{"pid": "1234"}

	h.Show(ctx1)
	h.Show(ctx2)
	h.Show(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestProductHandlerCreate(c *check.C) {
	h := NewProductHandler(s.db)

	ctx1, res1 := s.Context("POST", "/products", strings.NewReader(`{
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
	}`))
	ctx2, res2 := s.Context("POST", "/products", strings.NewReader(`{
		"name": "No Socks",
		"pricing": {
			"retail": -299
		}
	}`))

	h.Create(ctx1)
	h.Create(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestProductHandlerUpdate(c *check.C) {
	h := NewProductHandler(s.db)

	ctx1, res1 := s.Context("PUT", "/", strings.NewReader(`{
		"id": "`+product.Id.Hex()+`",
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
	}`))
	ctx1.Params = map[string]string{"pid": product.Id.Hex()}
	ctx2, res2 := s.Context("PUT", "/", strings.NewReader(`{
		"id": "`+product.Id.Hex()+`",
		"name": "Winter Socks",
		"pricing": {
			"retail": "599"
		},
		"description": "Getting cold again. Not with these socks anymore."
	}`))
	ctx2.Params = map[string]string{"pid": product.Id.Hex()}

	h.Update(ctx1)
	h.Update(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestProductHandlerDestroy(c *check.C) {
	h := NewProductHandler(s.db)

	ctx1, res1 := s.Context("DELETE", "/", nil)
	ctx1.Params = map[string]string{"pid": product.Id.Hex()}
	ctx2, res2 := s.Context("DELETE", "/", nil)
	ctx2.Params = map[string]string{"pid": bson.NewObjectId().Hex()}

	h.Destroy(ctx1)
	h.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}
