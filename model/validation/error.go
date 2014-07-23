package validation

import "fmt"

type Error map[string]error

func (err Error) Has() bool {
	return len(err) > 0
}

func (err Error) Set(k string, e error) {
	if err != nil {
		err[k] = e
	}
}

func (err Error) Error() string {
	for k, err := range err {
		return fmt.Sprintf("%s has %s", k, err)
	}
	return ""
}
