package context

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/check.v1"
)

func TestContext(t *testing.T) {
	check.Suite(&ContextSuite{})
	check.TestingT(t)
}

var (
	ErrTest = errors.New("test error")
)

type ContextSuite struct {
	context  *Context
	params   map[string]string
	request  *http.Request
	response *httptest.ResponseRecorder
}

func (s *ContextSuite) TestSet(c *check.C) {
	s.context.Set("X-Response-Time", "10ms")

	c.Check(s.response.Header().Get("X-Response-Time"), check.Equals, "10ms")
}

func (s *ContextSuite) TestGet(c *check.C) {
	c.Check(s.context.Get("Content-Type"), check.Equals, "application/json")
}

func (s *ContextSuite) TestQuery(c *check.C) {
	c.Check(s.context.Query("foo"), check.Equals, "bar")
}

func (s *ContextSuite) TestParam(c *check.C) {
	c.Check(s.context.Param("id"), check.Equals, "123")
}

func (s *ContextSuite) TestParse(c *check.C) {
	m := make(map[string]string)

	c.Check(s.context.Parse(&m), check.Equals, true)
	c.Check(m["foo"], check.Equals, "bar")
}

func (s *ContextSuite) TestErrorOne(c *check.C) {
	s.context.Error(nil, http.StatusNotFound)

	c.Check(s.response.Code, check.Equals, http.StatusNotFound)
	c.Check(s.response.Body.String(), check.Equals, "{\"error\":\"Not Found\"}\n")
}

func (s *ContextSuite) TestErrorCustom(c *check.C) {
	s.context.Error(ErrTest, http.StatusBadRequest)

	c.Check(s.response.Code, check.Equals, http.StatusBadRequest)
	c.Check(s.response.Body.String(), check.Equals, "{\"error\":\"test error\"}\n")
}

func (s *ContextSuite) TestRespond(c *check.C) {
	s.context.Respond(map[string]string{"foo": "bar"}, http.StatusBadRequest)

	c.Check(s.response.Code, check.Equals, http.StatusBadRequest)
	c.Check(s.response.Body.String(), check.Equals, "{\"foo\":\"bar\"}\n")
}

func (s *ContextSuite) SetUpTest(c *check.C) {
	s.response = httptest.NewRecorder()

	s.request, _ = http.NewRequest("POST", "/?foo=bar", strings.NewReader(`
		{"foo":"bar"}
	`))
	s.request.Header.Set("Content-Type", "application/json")

	s.context = &Context{
		Params: map[string]string{
			"id": "123",
		},
		Request:  s.request,
		Response: s.response,
	}
}
