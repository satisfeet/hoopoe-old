package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/check.v1"
)

func TestAuth(t *testing.T) {
	check.Suite(&AuthSuite{
		auth: &Auth{
			Username: "foo",
			Password: "bar",
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("Secure..."))
			}),
		},
	})
	check.TestingT(t)
}

type AuthSuite struct {
	auth *Auth
}

// Tests if everything works as expected with valid credentials.
func (s *AuthSuite) TestValid(c *check.C) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("foo", "bar")
	res := httptest.NewRecorder()

	s.auth.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusOK)
	c.Check(res.Body.String(), check.Equals, "Secure...")
}

// Tests if everything works as expected with invalid credentials.
func (s *AuthSuite) TestInvalidAuth(c *check.C) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.SetBasicAuth("foo", "baz")
	res := httptest.NewRecorder()

	s.auth.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusUnauthorized)
}

// Tests if no panic occurs on invalid header.
func (s *AuthSuite) TestInvalidHeader(c *check.C) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Basic ")
	res := httptest.NewRecorder()

	s.auth.ServeHTTP(res, req)

	c.Check(res.Code, check.Equals, http.StatusUnauthorized)
}
