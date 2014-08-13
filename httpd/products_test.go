package httpd

import (
	"net/http"
	"strings"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var product = model.Product{
	Name:   "Summer Socks",
	Images: []bson.ObjectId{},
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
	h := NewProductHandler(s.mongo)

	ctx1, res1 := s.Context("GET", "/products", nil)

	h.List(ctx1)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(strings.HasPrefix(res1.Body.String(), "[{"), check.Equals, true)
}

func (s *Suite) TestProductHandlerShow(c *check.C) {
	h := NewProductHandler(s.mongo)

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
	h := NewProductHandler(s.mongo)

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
	h := NewProductHandler(s.mongo)

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
	h := NewProductHandler(s.mongo)

	ctx1, res1 := s.Context("DELETE", "/", nil)
	ctx1.Params = map[string]string{"pid": product.Id.Hex()}
	ctx2, res2 := s.Context("DELETE", "/", nil)
	ctx2.Params = map[string]string{"pid": bson.NewObjectId().Hex()}

	h.Destroy(ctx1)
	h.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}

func (s *Suite) TestProductHandlerShowImage(c *check.C) {
	h := NewProductHandler(s.mongo)

	ctx1, res1 := s.Context("GET", "/", nil)
	ctx1.Params = map[string]string{
		"pid": product.Id.Hex(),
		"iid": product.Images[0].Hex(),
	}
	ctx2, res2 := s.Context("GET", "/", nil)
	ctx2.Params = map[string]string{
		"pid": "1234",
		"iid": product.Images[0].Hex(),
	}
	ctx3, res3 := s.Context("GET", "/", nil)
	ctx3.Params = map[string]string{
		"pid": product.Id.Hex(),
		"iid": "1234",
	}
	ctx4, res4 := s.Context("GET", "/", nil)
	ctx4.Params = map[string]string{
		"pid": product.Id.Hex(),
		"iid": bson.NewObjectId().Hex(),
	}

	h.ShowImage(ctx1)
	h.ShowImage(ctx2)
	h.ShowImage(ctx3)
	h.ShowImage(ctx4)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
	c.Check(res4.Code, check.Equals, http.StatusNotFound)

	c.Check(res1.Body.String(), check.Equals, "Hello World")
}

func (s *Suite) TestProductHandlerCreateImage(c *check.C) {
	h := NewProductHandler(s.mongo)

	ctx1, res1 := s.Context("POST", "/", strings.NewReader("Foo"))
	ctx1.Params = map[string]string{"pid": product.Id.Hex()}
	ctx2, res2 := s.Context("POST", "/", strings.NewReader("Foo"))
	ctx2.Params = map[string]string{"pid": bson.NewObjectId().Hex()}
	ctx3, res3 := s.Context("POST", "/", strings.NewReader("Foo"))
	ctx3.Params = map[string]string{"pid": "1234"}

	h.CreateImage(ctx1)
	h.CreateImage(ctx2)
	h.CreateImage(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestProductHandlerDestroyImage(c *check.C) {
	h := NewProductHandler(s.mongo)

	ctx1, res1 := s.Context("DELETE", "/", nil)
	ctx1.Params = map[string]string{
		"pid": product.Id.Hex(),
		"iid": product.Images[0].Hex(),
	}
	ctx2, res2 := s.Context("DELETE", "/", nil)
	ctx2.Params = map[string]string{
		"pid": "1234",
		"iid": product.Images[0].Hex(),
	}
	ctx3, res3 := s.Context("DELETE", "/", nil)
	ctx3.Params = map[string]string{
		"pid": product.Id.Hex(),
		"iid": "1234",
	}
	ctx4, res4 := s.Context("DELETE", "/", nil)
	ctx4.Params = map[string]string{
		"pid": product.Id.Hex(),
		"iid": bson.NewObjectId().Hex(),
	}

	h.DestroyImage(ctx1)
	h.DestroyImage(ctx2)
	h.DestroyImage(ctx3)
	h.DestroyImage(ctx4)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
	c.Check(res4.Code, check.Equals, http.StatusNotFound)

}
