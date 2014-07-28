package conf

import (
	"errors"
	"flag"
)

var (
	// Configuration values.
	Host     string
	Mongo    string
	Username string
	Password string

	// Possible configuration errors.
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
func Check() error {
	if len(Username) == 0 {
		return ErrUserInvalid
	}
	if len(Password) == 0 {
		return ErrPassInvalid
	}
	if len(Host) == 0 {
		return ErrHostInvalid
	}
	if len(Mongo) == 0 {
		return ErrMongoInvalid
	}

	return nil
}

// Flags loads data into Conf by parsing the provided
// arguments. The provided arguments are most likely to
// be os.Args[1:] but can also come from other sources.
func Flags(a []string) error {
	f := flag.NewFlagSet("conf", flag.ExitOnError)
	f.StringVar(&Username, "username", "", "Username for HTTP Basic.")
	f.StringVar(&Password, "password", "", "Password for HTTP Basic.")
	f.StringVar(&Host, "host", "", "Host address for HTTP Server.")
	f.StringVar(&Mongo, "mongo", "", "MongoDB URL for storage layer.")

	if err := f.Parse(a); err != nil {
		return err
	}

	return Check()
}
