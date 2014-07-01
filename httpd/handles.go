package httpd

import (
	"encoding/base64"
	"errors"
	"log"
	"strings"

	"github.com/satisfeet/hoopoe/httpd/router"
)

func Auth(c *router.Context) {
	h := strings.Split(c.Get("Authorization"), " ")

	if len(h) != 2 {
		c.RespondError(errors.New("Unauthorized"), 401)

		return
	}

	b, err := base64.StdEncoding.DecodeString(h[1])

	if err != nil {
		c.RespondError(errors.New("Bad Request"), 400)

		return
	}

	p := strings.Split(string(b), ":")

	if p[0] != "bodokaiser" {
		c.RespondError(errors.New("Unauthorized"), 401)

		return
	}

	c.Next()
}

func Logger(c *router.Context) {
	log.Printf("Request: %s %s", c.Method(), c.Path())

	c.Next()
}
