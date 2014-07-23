package conf

import (
	"errors"
	"flag"
)

// Conf contains application settings and
// methods to load settings from different sources.
type Conf struct {
	// Credentials used for HTTP Basic Auth.
	Username string
	Password string

	// Host address used from HTTP server to listen.
	Host string

	// Database servers storage layer will connect to.
	Mongo string
}

var (
	ErrUserInvalid  = errors.New("user parameter invalid")
	ErrPassInvalid  = errors.New("pass parameter invalid")
	ErrHostInvalid  = errors.New("host parameter invalid")
	ErrMongoInvalid = errors.New("mongo parameter invalid")
)

// Check validates the current values of Conf and
// returns an error if something is empty or invalid.
//
// Note that checks are still very basic and they will
// not guarante failures in other components.
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

// Flags loads data into Conf by parsing the provided
// arguments. The provided arguments are most likely to
// be os.Args[1:] but can also come from other sources.
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
