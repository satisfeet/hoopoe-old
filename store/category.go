package store

import "strings"

type Categories []string

func (c *Categories) Scan(src interface{}) error {
	switch t := src.(type) {
	case []byte:
		*c = strings.Split(string(t), ",")

		return nil
	}

	return ErrBadScanType
}
