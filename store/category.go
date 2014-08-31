package store

import (
	"errors"
	"strings"
)

var ErrBadCategoryType = errors.New("bad category type")
var ErrBadCategoryFormat = errors.New("bad category format")

type Categories []string

func (c *Categories) Scan(src interface{}) error {
	switch t := src.(type) {
	case []byte:
		*c = strings.Split(string(t), ",")

		return nil
	}

	return ErrBadCategoryType
}
