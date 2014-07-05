package conf

import (
	"flag"
)

const (
	ADDR  = "localhost:3000"
	MONGO = "localhost/satisfeet"
)

type Conf struct {
	Httpd Map
	Store Map
}

type Map map[string]string

func New() *Conf {
	return &Conf{
		Map{"addr": *flag.String("addr", ADDR, "HTTP address to listen.")},
		Map{"mongo": *flag.String("mongo", MONGO, "Mongo URL to connect.")},
	}
}
