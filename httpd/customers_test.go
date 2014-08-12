package httpd

import (
	"net/http"
	"strings"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store"
)

var customer = store.Customer{
	Name:  "Bob Marley",
	Email: "bob@yahoo.com",
	Address: store.Address{
		City: "Honolulu",
	},
}

func (s *Suite) TestCustomerHandlerList(c *check.C) {
	h := NewCustomerHandler(s.mongo)

	ctx1, res1 := s.Context("GET", "/customers", nil)
	ctx2, res2 := s.Context("GET", "/customers?search=mar", nil)
	ctx3, res3 := s.Context("GET", "/customers?search=foobar", nil)

	h.List(ctx1)
	h.List(ctx2)
	h.List(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusOK)
	c.Check(res3.Code, check.Equals, http.StatusOK)

	c.Check(strings.HasPrefix(res1.Body.String(), "[{"), check.Equals, true)
	c.Check(strings.HasPrefix(res2.Body.String(), "[{"), check.Equals, true)
	c.Check(strings.HasPrefix(res3.Body.String(), "[]"), check.Equals, true)
}

func (s *Suite) TestCustomerHandlerShow(c *check.C) {
	h := NewCustomerHandler(s.mongo)

	ctx1, res1 := s.Context("GET", "/", nil)
	ctx1.Params = map[string]string{"cid": customer.Id.Hex()}
	ctx2, res2 := s.Context("GET", "/", nil)
	ctx2.Params = map[string]string{"cid": bson.NewObjectId().Hex()}
	ctx3, res3 := s.Context("GET", "/", nil)
	ctx3.Params = map[string]string{"cid": "1234"}

	h.Show(ctx1)
	h.Show(ctx2)
	h.Show(ctx3)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
	c.Check(res3.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestCustomerHandlerCreate(c *check.C) {
	h := NewCustomerHandler(s.mongo)

	ctx1, res1 := s.Context("POST", "/customers", strings.NewReader(`{
		"name": "Edison T.",
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))
	ctx2, res2 := s.Context("POST", "/customers", strings.NewReader(`{
		"email": "edison@t.com",
		"address": {
			"city": "Leeds"
		}
	}`))

	h.Create(ctx1)
	h.Create(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusOK)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestCustomerHandlerUpdate(c *check.C) {
	h := NewCustomerHandler(s.mongo)

	ctx1, res1 := s.Context("PUT", "/", strings.NewReader(`{
		"id": "`+customer.Id.Hex()+`",
		"name": "Bob Marley",
		"email": "bob@marley.com",
		"address": {
			"city": "New York"
		}
	}`))
	ctx1.Params = map[string]string{"cid": customer.Id.Hex()}
	ctx2, res2 := s.Context("PUT", "/", strings.NewReader(`{
		"id": "`+customer.Id.Hex()+`",
		"name": "Bob Marley",
		"address": {
			"city": "New York"
		}
	}`))
	ctx2.Params = map[string]string{"cid": customer.Id.Hex()}

	h.Update(ctx1)
	h.Update(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusBadRequest)
}

func (s *Suite) TestCustomerHandlerDestroy(c *check.C) {
	h := NewCustomerHandler(s.mongo)

	ctx1, res1 := s.Context("DELETE", "/", nil)
	ctx1.Params = map[string]string{"cid": customer.Id.Hex()}
	ctx2, res2 := s.Context("DELETE", "/", nil)
	ctx2.Params = map[string]string{"cid": bson.NewObjectId().Hex()}

	h.Destroy(ctx1)
	h.Destroy(ctx2)

	c.Check(res1.Code, check.Equals, http.StatusNoContent)
	c.Check(res2.Code, check.Equals, http.StatusNotFound)
}
