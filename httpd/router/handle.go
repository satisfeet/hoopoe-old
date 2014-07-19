package router

import "github.com/satisfeet/hoopoe/httpd/context"

type Handler interface {
	ServeHTTP(*context.Context)
}

type HandlerFunc func(*context.Context)

func (handler HandlerFunc) ServeHTTP(c *context.Context) {
	handler(c)
}
