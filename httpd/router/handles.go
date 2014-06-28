package router

import (
	"log"
	"errors"
	"strings"
	"encoding/base64"
)

func Auth(c *Context) {
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

func Logger(c *Context) {
	log.Printf("Request: %s %s", c.Method(), c.Path())

	c.Next()
}
