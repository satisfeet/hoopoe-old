package httpd

import (
	"net/http"
	"strings"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

var customer = model.Customer{
	Id:    bson.NewObjectId(),
	Name:  "Bob Marley",
	Email: "bob@yahoo.com",
	Address: model.Address{
		City: "Honolulu",
	},
}

var cs = &CustomerSuite{
	HandlerSuite: hs,
}

func TestCustomer(t *testing.T) {
	check.Suite(hs)
	check.TestingT(t)
}

type CustomerSuite struct {
	*HandlerSuite
	handler *Customer
}

func (s *CustomerSuite) SetUpTest(c *check.C) {
	s.handler = NewCustomer(s.session)

	err := s.database.C("customers").Insert(customer)
	c.Assert(err, check.IsNil)
}

func (s *CustomerSuite) TestList(c *check.C) {
	ctx1, res1 := ctx("GET", "/customers", nil)
	ctx2, res2 := ctx("GET", "/customers?search=mar", nil)
	ctx3, res3 := ctx("GET", "/customers?search=foobar", nil)

	s.handler.List(ctx1)
	s.handler.List(ctx2)
	s.handler.List(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusOK)
	c.Check(res3.Code, check.Equals, http.StatusOK)

	c.Check(strings.HasPrefix(res1.Body.String(), "[{"), check.Equals, true)
	c.Check(strings.HasPrefix(res2.Body.String(), "[{"), check.Equals, true)
	c.Check(strings.HasPrefix(res3.Body.String(), "[]"), check.Equals, true)
}

func (s *CustomerSuite) TestShow(c *check.C) {
	ctx1, res1 := ctx("GET", "/", nil)
	ctx2, res2 := ctx("GET", "/", nil)
	ctx3, res3 := ctx("GET", "/", nil)

	ctx1.Params = map[string]string{"customer": customer.Id.Hex()}
	ctx2.Params = map[string]string{"customer": bson.NewObjectId().Hex()}
	ctx3.Params = map[string]string{"customer": "1234"}

	s.handler.Show(ctx1)
	s.handler.Show(ctx2)
	s.handler.Show(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *CustomerSuite) TestCreate(c *check.C) {
	ctx1, res1 := ctx("POST", "/customers", strings.NewReader(`{
		"name": "Edison T.",
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))
	ctx2, res2 := ctx("POST", "/customers", strings.NewReader(`{
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))

	s.handler.Create(ctx1)
	s.handler.Create(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *CustomerSuite) TestUpdate(c *check.C) {
	ctx1, res1 := ctx("PUT", "/", strings.NewReader(`{
		"id": "`+customer.Id.Hex()+`",
		"name": "Bob Marley",
		"email": "bob@marley.com",
		"address": {
			"city": "New York"
		}
	}`))
	ctx2, res2 := ctx("PUT", "/", strings.NewReader(`{
		"id": "`+customer.Id.Hex()+`",
		"name": "Bob Marley",
		"address": {
			"city": "New York"
		}
	}`))

	ctx1.Params = map[string]string{"customer": customer.Id.Hex()}
	ctx2.Params = map[string]string{"customer": customer.Id.Hex()}

	s.handler.Update(ctx1)
	s.handler.Update(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *CustomerSuite) TestDestroy(c *check.C) {
	ctx1, res1 := ctx("DELETE", "/", nil)
	ctx2, res2 := ctx("DELETE", "/", nil)

	ctx1.Params = map[string]string{"customer": customer.Id.Hex()}
	ctx2.Params = map[string]string{"customer": bson.NewObjectId().Hex()}

	s.handler.Destroy(ctx1)
	s.handler.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}

func (s *CustomerSuite) TearDownTest(c *check.C) {
	_, err := s.database.C("customers").RemoveAll(nil)
	c.Assert(err, check.IsNil)
}
