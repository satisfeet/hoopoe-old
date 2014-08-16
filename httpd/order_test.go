package httpd

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store"
)

var order = model.Order{
	Id:     bson.NewObjectId(),
	Number: 1,
	Items: []model.OrderItem{
		model.OrderItem{
			ProductId: product.Id,
			Pricing:   product.Pricing,
			Quantity:  1,
		},
	},
	Pricing:    product.Pricing,
	CustomerId: customer.Id,
}

func TestOrder(t *testing.T) {
	hs := &HandlerSuite{
		url: "localhost/test",
	}

	check.Suite(&OrderSuite{
		HandlerSuite:  hs,
		ProductSuite:  &ProductSuite{HandlerSuite: hs},
		CustomerSuite: &CustomerSuite{HandlerSuite: hs},
	})
	check.TestingT(t)
}

type OrderSuite struct {
	handler *Order
	*HandlerSuite
	ProductSuite  *ProductSuite
	CustomerSuite *CustomerSuite
}

func (s *OrderSuite) SetUpSuite(c *check.C) {
	s.HandlerSuite.SetUpSuite(c)
	s.ProductSuite.SetUpSuite(c)
	s.CustomerSuite.SetUpSuite(c)
}

func (s *OrderSuite) SetUpTest(c *check.C) {
	s.ProductSuite.SetUpTest(c)
	s.CustomerSuite.SetUpTest(c)

	err := s.database.C("orders").EnsureIndex(mgo.Index{
		Key:    store.OrderUnique,
		Unique: true,
	})
	c.Assert(err, check.IsNil)

	err = s.database.C("orders").Insert(order)
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("orders").Create("")
	f.SetId(order.Id)
	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Invoice"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)

	s.handler = NewOrder(s.session)
}

func (s *OrderSuite) TestList(c *check.C) {
	ctx, res := ctx("GET", "/orders", nil)

	s.handler.List(ctx)

	o := []model.Order{}

	err := json.NewDecoder(res.Body).Decode(&o)
	c.Assert(err, check.IsNil)

	c.Check(res.Code, check.Equals, http.StatusOK)

	c.Check(o, check.HasLen, 1)
	c.Check(o[0], check.DeepEquals, order)
}

func (s *OrderSuite) TestShow(c *check.C) {
	ctx1, res1 := ctx("GET", "/", nil)
	ctx2, res2 := ctx("GET", "/", nil)
	ctx3, res3 := ctx("GET", "/", nil)

	ctx1.Params = map[string]string{"order": order.Id.Hex()}
	ctx2.Params = map[string]string{"order": bson.NewObjectId().Hex()}
	ctx3.Params = map[string]string{"order": "1234"}

	s.handler.Show(ctx1)
	s.handler.Show(ctx2)
	s.handler.Show(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *OrderSuite) TestCreate(c *check.C) {
	ctx1, res1 := ctx("POST", "/orders", strings.NewReader(`{
		"items": [
			{
				"product": "`+product.Id.Hex()+`",
				"quantity": 1,
				"pricing": {
					"retail": 299
				},
				"variation": {
					"size": "`+product.Variations[0].Size+`",
					"color": "`+product.Variations[0].Color+`"
				}
			}
		],
		"customer": "`+customer.Id.Hex()+`",
		"pricing": {
			"retail": 299
		}
	}`))
	ctx2, res2 := ctx("POST", "/orders", strings.NewReader(`{
		"items": [
			{
				"product": "`+product.Id.Hex()+`",
				"quantity": 1,
				"pricing": {
					"retail": 299
				},
				"variation": {
					"size": "`+product.Variations[0].Size+`",
					"color": "`+product.Variations[0].Color+`"
				}
			}
		],
		"pricing": {
			"retail": 299
		}
	}`))

	s.handler.Create(ctx1)
	s.handler.Create(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *OrderSuite) TestDestroy(c *check.C) {
	ctx1, res1 := ctx("DELETE", "/", nil)
	ctx2, res2 := ctx("DELETE", "/", nil)

	ctx1.Params = map[string]string{"order": order.Id.Hex()}
	ctx2.Params = map[string]string{"order": bson.NewObjectId().Hex()}

	s.handler.Destroy(ctx1)
	s.handler.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}

func (s *OrderSuite) TestShowInvoice(c *check.C) {
	ctx1, res1 := ctx("GET", "/", nil)
	ctx2, res2 := ctx("GET", "/", nil)
	ctx3, res3 := ctx("GET", "/", nil)

	ctx1.Params = map[string]string{"order": order.Id.Hex()}
	ctx2.Params = map[string]string{"order": bson.NewObjectId().Hex()}
	ctx3.Params = map[string]string{"order": "12345"}

	s.handler.ShowInvoice(ctx1)
	s.handler.ShowInvoice(ctx2)
	s.handler.ShowInvoice(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)

	c.Check(res1.Body.String(), check.Equals, "Invoice")
}

func (s *OrderSuite) TearDownTest(c *check.C) {
	s.ProductSuite.TearDownTest(c)
	s.CustomerSuite.TearDownTest(c)

	err := s.database.C("orders").DropIndex(store.OrderUnique...)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("orders").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("orders.files").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("orders.chunks").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TearDownSuite(c *check.C) {
	s.CustomerSuite.TearDownSuite(c)
	s.ProductSuite.TearDownSuite(c)
	s.HandlerSuite.TearDownSuite(c)
}
