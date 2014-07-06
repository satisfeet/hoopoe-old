package conf

import (
	"errors"
	"flag"
)

type Map map[string]string

type Conf struct {
	Store Map
	Httpd Map
}

func New() *Conf {
	return &Conf{
		Store: Map{},
		Httpd: Map{},
	}
}

func (c *Conf) ParseFlags() error {
	a := flag.String("addr", "", "HTTP address to listen.")
	m := flag.String("mongo", "", "MongoDB URL to connect.")

	flag.Parse()

	if len(*a) == 0 {
		return errors.New(`"addr" flag is required.`)
	}
	if len(*m) == 0 {
		return errors.New(`"mongo" flag is required.`)
	}

	c.Store["mongo"] = *m
	c.Httpd["addr"] = *a

	return nil
}
