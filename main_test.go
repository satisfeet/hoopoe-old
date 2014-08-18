package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
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
	url     string
	user    string
	pass    string
	session *mgo.Session
}

func (s *Suite) SetUpSuite(c *check.C) {
	sess, err := mgo.Dial(s.url)
	c.Assert(err, check.IsNil)

	s.session = sess

	auth = s.user + ":" + s.pass
}

func (s *Suite) TestHandler(c *check.C) {
	h := Handler(s.session)

	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()
	res3 := httptest.NewRecorder()
	res4 := httptest.NewRecorder()
	res5 := httptest.NewRecorder()

	req1, _ := http.NewRequest("GET", "/", nil)
	req2, _ := http.NewRequest("GET", "/not-found", nil)
	req3, _ := http.NewRequest("GET", "/products", nil)
	req4, _ := http.NewRequest("GET", "/customers", nil)
	req5, _ := http.NewRequest("GET", "/orders", nil)

	req2.SetBasicAuth(s.user, s.pass)
	req3.SetBasicAuth(s.user, s.pass)
	req4.SetBasicAuth(s.user, s.pass)
	req5.SetBasicAuth(s.user, s.pass)

	h.ServeHTTP(res1, req1)
	h.ServeHTTP(res2, req2)
	h.ServeHTTP(res3, req3)
	h.ServeHTTP(res4, req4)
	h.ServeHTTP(res5, req5)
}
