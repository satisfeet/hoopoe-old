package httpd

import (
	"net/http"
	"strings"
	"testing"

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

var ps = &ProductSuite{
	HandlerSuite: hs,
}

func TestProduct(t *testing.T) {
	check.Suite(ps)
	check.TestingT(t)
}

type ProductSuite struct {
	*HandlerSuite
	handler *Product
}

func (s *ProductSuite) SetUpTest(c *check.C) {
	s.handler = NewProduct(s.mongo)

	err := s.mongo.Insert("products", &product)
	c.Assert(err, check.IsNil)

	f, err := s.mongo.CreateFile("products")
	c.Assert(err, check.IsNil)
	product.Images = []bson.ObjectId{f.Id}
	_, err = f.Write([]byte("Hello World"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TestList(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("GET", "/products", nil)

	h.List(ctx1)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(strings.HasPrefix(res1.Body.String(), "[{"), check.Equals, true)
}

func (s *ProductSuite) TestShow(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("GET", "/", nil)
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("GET", "/", nil)
	ctx2.Params = map[string]string{"product": bson.NewObjectId().Hex()}
	ctx3, res3 := ctx("GET", "/", nil)
	ctx3.Params = map[string]string{"product": "1234"}

	h.Show(ctx1)
	h.Show(ctx2)
	h.Show(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *ProductSuite) TestCreate(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("POST", "/products", strings.NewReader(`{
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
	ctx2, res2 := ctx("POST", "/products", strings.NewReader(`{
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

func (s *ProductSuite) TestUpdate(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("PUT", "/", strings.NewReader(`{
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
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("PUT", "/", strings.NewReader(`{
		"id": "`+product.Id.Hex()+`",
		"name": "Winter Socks",
		"pricing": {
			"retail": "599"
		},
		"description": "Getting cold again. Not with these socks anymore."
	}`))
	ctx2.Params = map[string]string{"product": product.Id.Hex()}

	h.Update(ctx1)
	h.Update(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *ProductSuite) TestDestroy(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("DELETE", "/", nil)
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("DELETE", "/", nil)
	ctx2.Params = map[string]string{"product": bson.NewObjectId().Hex()}

	h.Destroy(ctx1)
	h.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}

func (s *ProductSuite) TestShowImage(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("GET", "/", nil)
	ctx1.Params = map[string]string{
		"product": product.Id.Hex(),
		"image":   product.Images[0].Hex(),
	}
	ctx2, res2 := ctx("GET", "/", nil)
	ctx2.Params = map[string]string{
		"product": "1234",
		"image":   product.Images[0].Hex(),
	}
	ctx3, res3 := ctx("GET", "/", nil)
	ctx3.Params = map[string]string{
		"product": product.Id.Hex(),
		"image":   "1234",
	}
	ctx4, res4 := ctx("GET", "/", nil)
	ctx4.Params = map[string]string{
		"product": product.Id.Hex(),
		"image":   bson.NewObjectId().Hex(),
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

func (s *ProductSuite) TestCreateImage(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("POST", "/", strings.NewReader("Foo"))
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("POST", "/", strings.NewReader("Foo"))
	ctx2.Params = map[string]string{"product": bson.NewObjectId().Hex()}
	ctx3, res3 := ctx("POST", "/", strings.NewReader("Foo"))
	ctx3.Params = map[string]string{"product": "1234"}

	h.CreateImage(ctx1)
	h.CreateImage(ctx2)
	h.CreateImage(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *ProductSuite) TestDestroyImage(c *check.C) {
	h := NewProduct(s.mongo)

	ctx1, res1 := ctx("DELETE", "/", nil)
	ctx1.Params = map[string]string{
		"product": product.Id.Hex(),
		"image":   product.Images[0].Hex(),
	}
	ctx2, res2 := ctx("DELETE", "/", nil)
	ctx2.Params = map[string]string{
		"product": "1234",
		"image":   product.Images[0].Hex(),
	}
	ctx3, res3 := ctx("DELETE", "/", nil)
	ctx3.Params = map[string]string{
		"product": product.Id.Hex(),
		"image":   "1234",
	}
	ctx4, res4 := ctx("DELETE", "/", nil)
	ctx4.Params = map[string]string{
		"product": product.Id.Hex(),
		"image":   bson.NewObjectId().Hex(),
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

func (s *ProductSuite) TearDownTest(c *check.C) {
	c.Assert(s.mongo.RemoveAll("products", nil), check.IsNil)
	c.Assert(s.mongo.RemoveAll("products.files", nil), check.IsNil)
	c.Assert(s.mongo.RemoveAll("products.chunks", nil), check.IsNil)
}
