package conf

import (
	"errors"
	"flag"
)

type Conf struct {
	Host     string
	Mongo    string
	Username string
	Password string
}

func NewConf() *Conf {
	return &Conf{}
}

var (
	ErrUserInvalid  = errors.New("user parameter invalid")
	ErrPassInvalid  = errors.New("pass parameter invalid")
	ErrHostInvalid  = errors.New("host parameter invalid")
	ErrMongoInvalid = errors.New("mongo parameter invalid")
)

func (c *Conf) Check() error {
	if len(c.Username) == 0 {
		return ErrUserInvalid
	}
	if len(c.Password) == 0 {
		return ErrPassInvalid
	}
	if len(c.Host) == 0 {
		return ErrHostInvalid
	}
	if len(c.Mongo) == 0 {
		return ErrMongoInvalid
	}

	return nil
}

func (c *Conf) Flags(a []string) error {
	f := flag.NewFlagSet("conf", flag.ExitOnError)
	f.StringVar(&c.Username, "username", "", "Username for HTTP Basic.")
	f.StringVar(&c.Password, "password", "", "Password for HTTP Basic.")
	f.StringVar(&c.Host, "host", "", "Host address for HTTP Server.")
	f.StringVar(&c.Mongo, "mongo", "", "MongoDB URL for storage layer.")

	if err := f.Parse(a); err != nil {
		return err
	}

	return c.Check()
}
