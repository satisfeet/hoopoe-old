package store

import (
	"io/ioutil"
	"testing"
	"time"

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
	State: model.OrderState{
		Created: time.Date(2013, time.November, 10, 23, 0, 0, 0, time.Local),
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

	err := s.database.C("orders").EnsureIndex(mgo.Index{
		Key:    OrderUnique,
		Unique: true,
	})
	c.Assert(err, check.IsNil)

	err = s.database.C("orders").Insert(order)
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("orders").Create("")
	f.SetId(order.Id)

	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TestIndex(c *check.C) {
	err := s.store.Index()
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

func (s *OrderSuite) TestInsert(c *check.C) {
	m1 := model.Order{
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
	m2 := model.Order{
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

	err := s.store.Insert(&m1)
	c.Assert(err, check.IsNil)
	err = s.store.Insert(&m2)
	c.Assert(err, check.IsNil)

	c.Check(m1.Number, check.Equals, 1)
	c.Check(m2.Number, check.Equals, 2)
}

func (s *OrderSuite) TestRemove(c *check.C) {
	err := s.store.Remove(order)
	c.Assert(err, check.IsNil)
}

func (s *OrderSuite) TestOpenInvoice(c *check.C) {
	rc, err := s.store.OpenInvoice(&order)
	c.Assert(err, check.IsNil)
	defer rc.Close()

	b, err := ioutil.ReadAll(rc)
	c.Assert(err, check.IsNil)
	c.Assert(b, check.DeepEquals, []byte("Hello"))
}

func (s *OrderSuite) TestCreateInvoice(c *check.C) {
	err := s.database.GridFS("orders").RemoveId(order.Id)
	c.Assert(err, check.IsNil)

	wc, err := s.store.CreateInvoice(&order)
	c.Assert(err, check.IsNil)
	_, err = wc.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = wc.Close()
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

	err := s.database.C("orders").DropIndex(OrderUnique...)
	c.Assert(err, check.IsNil)

	_, err = s.database.C("orders").RemoveAll(nil)
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
