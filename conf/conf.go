package conf

import (
	"flag"
)

const (
	ADDR  = "localhost:3000"
	MONGO = "localhost/satisfeet"
)

type Conf struct {
	Httpd map[string]string
	Store map[string]string
}

func New() *Conf {
	return &Conf{
		map[string]string{
			"addr": *flag.String("addr", ADDR, "HTTP address to listen."),
		},
		map[string]string{
			"mongo": *flag.String("mongo", MONGO, "Mongo URL to connect."),
		},
	}
}
