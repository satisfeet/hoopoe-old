package conf

import (
	"errors"
	"flag"
	"strings"
)

// Conf contains application settings and
// methods to load settings from different sources.
type Conf struct {
	// Credentials used for HTTP Basic Auth.
	Auth string

	// Host address used from HTTP server to listen.
	Host string

	// Database servers storage layer will connect to.
	Mongo string
}

var (
	ErrAuthInvalid = errors.New("auth parameter invalid")
	ErrHostInvalid = errors.New("host parameter invalid")

	ErrAuthRequired  = errors.New("auth parameter required")
	ErrHostRequired  = errors.New("host parameter required")
	ErrMongoRequired = errors.New("mongo parameter required")
)

// Check validates the current values of Conf and
// returns an error if something is empty or invalid.
//
// Note that checks are still very basic and they will
// not guarante failures in other components.
func (c *Conf) Check() error {
	if len(c.Auth) == 0 {
		return ErrAuthRequired
	}
	if len(c.Host) == 0 {
		return ErrHostRequired
	}
	if len(c.Mongo) == 0 {
		return ErrMongoRequired
	}

	if !strings.Contains(c.Auth, ":") {
		return ErrAuthInvalid
	}
	if !strings.Contains(c.Host, ":") {
		return ErrHostInvalid
	}

	return nil
}

// Flags loads data into Conf by parsing the provided
// arguments. The provided arguments are most likely to
// be os.Args[1:] but can also come from other sources.
func (c *Conf) Flags(a []string) error {
	f := flag.NewFlagSet("conf", flag.ExitOnError)
	f.StringVar(&c.Auth, "auth", "", "Auth credentials for HTTP Basic.")
	f.StringVar(&c.Host, "host", "", "Host address for HTTP Server.")
	f.StringVar(&c.Mongo, "mongo", "", "MongoDB URL for storage layer.")

	if err := f.Parse(a); err != nil {
		return err
	}

	return c.Check()
}
