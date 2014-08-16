package store

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var order = model.Order{
	Id: bson.NewObjectId(),
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
	check.Suite(&OrderSuite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type OrderSuite struct {
	url      string
	store    *Order
	session  *mgo.Session
	database *mgo.Database
	product  *ProductSuite
	customer *CustomerSuite
}

func (s *OrderSuite) SetUpSuite(c *check.C) {
	s.product = &ProductSuite{url: s.url}
	s.customer = &CustomerSuite{url: s.url}

	s.product.SetUpSuite(c)
	s.customer.SetUpSuite(c)

	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.session = sess
	s.database = sess.DB("")

	s.store = NewOrder(sess)
}

func (s *OrderSuite) SetUpTest(c *check.C) {
	s.product.SetUpTest(c)
	s.customer.SetUpTest(c)

	err := s.database.C("orders").Insert(order)
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("orders").Create("")
	f.SetId(order.Id)

	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TestFind(c *check.C) {
	m := []model.Order{}

	err := s.store.Find(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.HasLen, 1)
	c.Check(m[0], check.DeepEquals, order)
}

func (s *OrderSuite) TestFindOne(c *check.C) {
	m1 := model.Order{Id: order.Id}
	m2 := model.Order{Id: bson.ObjectId("1234")}

	err1 := s.store.FindOne(&m1)
	err2 := s.store.FindOne(&m2)

	c.Assert(err1, check.IsNil)
	c.Assert(err2, check.Equals, ErrBadId)

	c.Check(m1, check.DeepEquals, order)
}

func (s *OrderSuite) TestUpdate(c *check.C) {
	order.Pricing.Retail += 200

	err := s.store.Update(&order)
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TestRemove(c *check.C) {
	err := s.store.Remove(order)
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TestReadInvoice(c *check.C) {
	buf := &bytes.Buffer{}

	err := s.store.ReadInvoice(&order, buf)
	c.Assert(err, check.IsNil)

	c.Check(buf.String(), check.Equals, "Hello")
}

func (s *OrderSuite) TestWriteInvoice(c *check.C) {
	err := s.database.GridFS("orders").RemoveId(order.Id)
	c.Assert(err, check.IsNil)

	err = s.store.WriteInvoice(&order, strings.NewReader("Hello"))
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("orders").OpenId(order.Id)
	c.Assert(err, check.IsNil)
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	c.Assert(err, check.IsNil)

	c.Check(b, check.DeepEquals, []byte("Hello"))
}

func (s *OrderSuite) TearDownTest(c *check.C) {
	s.product.TearDownTest(c)
	s.customer.TearDownTest(c)

	_, err := s.database.C("orders").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("orders.files").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("orders.chunks").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TearDownSuite(c *check.C) {
	s.product.TearDownSuite(c)
	s.customer.TearDownSuite(c)

	s.session.Close()
}
