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
	Id:     bson.NewObjectId(),
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

func TestProduct(t *testing.T) {
	check.Suite(&ProductSuite{
		HandlerSuite: &HandlerSuite{
			url: "localhost/test",
		},
	})
	check.TestingT(t)
}

type ProductSuite struct {
	*HandlerSuite
	handler *Product
}

func (s *ProductSuite) SetUpTest(c *check.C) {
	s.handler = NewProduct(s.session)

	err := s.database.C("products").Insert(product)
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("products").Create("")
	c.Assert(err, check.IsNil)
	product.Images = []bson.ObjectId{f.Id().(bson.ObjectId)}
	_, err = f.Write([]byte("Hello World"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TestList(c *check.C) {
	ctx1, res1 := ctx("GET", "/products", nil)

	s.handler.List(ctx1)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(strings.HasPrefix(res1.Body.String(), "[{"), check.Equals, true)
}

func (s *ProductSuite) TestShow(c *check.C) {
	ctx1, res1 := ctx("GET", "/", nil)
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("GET", "/", nil)
	ctx2.Params = map[string]string{"product": bson.NewObjectId().Hex()}
	ctx3, res3 := ctx("GET", "/", nil)
	ctx3.Params = map[string]string{"product": "1234"}

	s.handler.Show(ctx1)
	s.handler.Show(ctx2)
	s.handler.Show(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *ProductSuite) TestCreate(c *check.C) {
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

	s.handler.Create(ctx1)
	s.handler.Create(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *ProductSuite) TestUpdate(c *check.C) {
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

	s.handler.Update(ctx1)
	s.handler.Update(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *ProductSuite) TestDestroy(c *check.C) {
	ctx1, res1 := ctx("DELETE", "/", nil)
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("DELETE", "/", nil)
	ctx2.Params = map[string]string{"product": bson.NewObjectId().Hex()}

	s.handler.Destroy(ctx1)
	s.handler.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}

func (s *ProductSuite) TestShowImage(c *check.C) {
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

	s.handler.ShowImage(ctx1)
	s.handler.ShowImage(ctx2)
	s.handler.ShowImage(ctx3)
	s.handler.ShowImage(ctx4)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
	c.Check(res4.Code, check.Equals, http.StatusNotFound)

	c.Check(res1.Body.String(), check.Equals, "Hello World")
}

func (s *ProductSuite) TestCreateImage(c *check.C) {
	ctx1, res1 := ctx("POST", "/", strings.NewReader("Foo"))
	ctx1.Params = map[string]string{"product": product.Id.Hex()}
	ctx2, res2 := ctx("POST", "/", strings.NewReader("Foo"))
	ctx2.Params = map[string]string{"product": bson.NewObjectId().Hex()}
	ctx3, res3 := ctx("POST", "/", strings.NewReader("Foo"))
	ctx3.Params = map[string]string{"product": "1234"}

	s.handler.CreateImage(ctx1)
	s.handler.CreateImage(ctx2)
	s.handler.CreateImage(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
	c.Logf("error nessage: %s\n", res3.Body.String())
}

func (s *ProductSuite) TestDestroyImage(c *check.C) {
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

	s.handler.DestroyImage(ctx1)
	s.handler.DestroyImage(ctx2)
	s.handler.DestroyImage(ctx3)
	s.handler.DestroyImage(ctx4)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
	c.Check(res4.Code, check.Equals, http.StatusNotFound)
}

func (s *ProductSuite) TearDownTest(c *check.C) {
	_, err := s.database.C("products").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("products.files").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("products.chunks").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}
