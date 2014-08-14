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

var product = model.Product{
	Id:   bson.NewObjectId(),
	Name: "Summer socks",
	Images: []bson.ObjectId{
		bson.NewObjectId(),
	},
	Pricing: model.Pricing{
		Retail: 599,
	},
	Variations: []model.Variation{
		model.Variation{
			Size:  "42-44",
			Color: "black",
		},
	},
	Description: "These are great socks.... I can really recommend them....",
}

func TestProduct(t *testing.T) {
	check.Suite(&ProductSuite{
		url: "localhost/test",
	})
	check.TestingT(t)
}

type ProductSuite struct {
	url      string
	store    *Product
	session  *mgo.Session
	database *mgo.Database
}

func (s *ProductSuite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.session = sess
	s.database = sess.DB("")

	s.store = NewProduct(sess)
}

func (s *ProductSuite) SetUpTest(c *check.C) {
	err := s.database.C("products").Insert(product)
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("products").Create("")
	f.SetId(product.Images[0])

	c.Assert(err, check.IsNil)
	_, err = f.Write([]byte("Hello"))
	c.Assert(err, check.IsNil)
	err = f.Close()
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TestFind(c *check.C) {
	m := []model.Product{}

	err := s.store.Find(&m)
	c.Assert(err, check.IsNil)

	c.Check(m, check.HasLen, 1)
	c.Check(m[0], check.DeepEquals, product)
}

func (s *ProductSuite) TestFindOne(c *check.C) {
	m1 := model.Product{Id: product.Id}
	m2 := model.Product{Id: bson.ObjectId("1234")}

	err1 := s.store.FindOne(&m1)
	err2 := s.store.FindOne(&m2)

	c.Assert(err1, check.IsNil)
	c.Assert(err2, check.Equals, ErrBadId)

	c.Check(m1, check.DeepEquals, product)
}

func (s *ProductSuite) TestInsert(c *check.C) {
	m := model.Product{
		Name:        "Winter Socks",
		Pricing:     product.Pricing,
		Variations:  product.Variations,
		Description: product.Description,
	}

	err := s.store.Insert(&m)
	c.Assert(err, check.IsNil)

	c.Check(m.Id.Valid(), check.Equals, true)
}

func (s *ProductSuite) TestUpdate(c *check.C) {
	product.Name += " Jr"

	err := s.store.Update(&product)
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TestRemove(c *check.C) {
	err := s.store.Remove(product)
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TestReadImage(c *check.C) {
	buf := &bytes.Buffer{}

	err := s.store.ReadImage(&product, product.Images[0], buf)
	c.Assert(err, check.IsNil)

	c.Check(buf.String(), check.Equals, "Hello")
}

func (s *ProductSuite) TestWriteImage(c *check.C) {
	id := bson.NewObjectId()
	buf := strings.NewReader("World")

	err := s.store.WriteImage(&product, id, buf)
	c.Assert(err, check.IsNil)

	f, err := s.database.GridFS("products").OpenId(id)
	c.Assert(err, check.IsNil)
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	c.Assert(err, check.IsNil)

	c.Check(b, check.DeepEquals, []byte("World"))
}

func (s *ProductSuite) TestRemoveImage(c *check.C) {
	err := s.store.RemoveImage(&product, product.Images[0])
	c.Assert(err, check.IsNil)

	_, err = s.database.GridFS("products").OpenId(product.Images[0])
	c.Assert(err, check.NotNil)
}

func (s *ProductSuite) TearDownTest(c *check.C) {
	_, err := s.database.C("products").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("products.files").RemoveAll(nil)
	c.Assert(err, check.IsNil)
	_, err = s.database.C("products.chunks").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}

func (s *ProductSuite) TearDownSuite(c *check.C) {
	s.session.Close()
}
