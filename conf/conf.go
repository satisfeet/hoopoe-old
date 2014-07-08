package conf

import (
	"errors"
	"flag"
)

type Conf map[string]string

func NewConf() Conf {
	return Conf{}
}

func (c Conf) ParseFlags() error {
	a := flag.String("addr", "", "HTTP address to listen.")
	m := flag.String("mongo", "", "MongoDB URL to connect.")

	flag.Parse()

	if len(*a) == 0 {
		return errors.New(`"addr" flag is required.`)
	}
	if len(*m) == 0 {
		return errors.New(`"mongo" flag is required.`)
	}

	c["mongo"] = *m
	c["addr"] = *a

	return nil
}
