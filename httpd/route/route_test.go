package route

import (
	"net/http"
	"testing"

	"gopkg.in/check.v1"
)

func TestRoute(t *testing.T) {
	check.Suite(&RouteSuite{
		Prefix: "/models",
	})
	check.TestingT(t)
}

type RouteSuite struct {
	Prefix string
}

func (s *RouteSuite) route(m, p string) (Action, string) {
	r, _ := http.NewRequest(m, p, nil)

	return Route(s.Prefix, r)
}

func (s *RouteSuite) Test(c *check.C) {
	a, p := s.route("GET", "/models")
	c.Check(a, check.Equals, List)
	c.Check(p, check.Equals, "")
	a, p = s.route("POST", "/models")
	c.Check(a, check.Equals, Create)
	c.Check(p, check.Equals, "")

	a, p = s.route("GET", "/models/1234")
	c.Check(a, check.Equals, Show)
	c.Check(p, check.Equals, "1234")
	a, p = s.route("PUT", "/models/1234")
	c.Check(a, check.Equals, Update)
	c.Check(p, check.Equals, "1234")

	a, p = s.route("GET", "/")
	c.Check(a, check.Equals, Invalid)
	a, p = s.route("GET", "/modelss")
	c.Check(a, check.Equals, Invalid)
	a, p = s.route("GET", "/models/12/34")
	c.Check(a, check.Equals, Invalid)
	a, p = s.route("GET", "/modelss/12/34")
	c.Check(a, check.Equals, Invalid)
}
