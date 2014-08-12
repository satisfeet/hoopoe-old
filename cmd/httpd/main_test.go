package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satisfeet/hoopoe/store/mongo"
	"gopkg.in/check.v1"
)

func TestMain(t *testing.T) {
	check.Suite(&Suite{
		url:  "localhost/test",
		user: "foo",
		pass: "bar",
	})
	check.TestingT(t)
}

type Suite struct {
	url   string
	user  string
	pass  string
	mongo *mongo.Store
}

func (s *Suite) SetUpSuite(c *check.C) {
	s.mongo = &mongo.Store{}
	c.Assert(s.mongo.Dial(s.url), check.IsNil)

	auth = s.user + ":" + s.pass
}

func (s *Suite) TestHandle(c *check.C) {
	h := Handle(s.mongo)

	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()
	res3 := httptest.NewRecorder()
	res4 := httptest.NewRecorder()

	req1, _ := http.NewRequest("GET", "/", nil)
	req2, _ := http.NewRequest("GET", "/not-found", nil)
	req3, _ := http.NewRequest("GET", "/products", nil)
	req4, _ := http.NewRequest("GET", "/customers", nil)

	req2.SetBasicAuth(s.user, s.pass)
	req3.SetBasicAuth(s.user, s.pass)
	req4.SetBasicAuth(s.user, s.pass)

	h.ServeHTTP(res1, req1)
	h.ServeHTTP(res2, req2)
	h.ServeHTTP(res3, req3)
	h.ServeHTTP(res4, req4)
}
