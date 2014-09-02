package mapper

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

type MapperFunc func(src string, target interface{}) error

var (
	ErrBadSource = errors.New("bad source type")
	ErrBadTarget = errors.New("bad target type")
)

func MapIntFunc(src string, target interface{}) error {
	var i int64
	var err error

	switch t := target.(type) {
	case *int:
		if len(src) != 0 {
			i, err = strconv.ParseInt(src, 0, 32)
		} else {
			i = 0
		}
		*t = int(i)

		return err
	case *int64:
		if len(src) != 0 {
			i, err = strconv.ParseInt(src, 0, 64)
		} else {
			i = 0
		}
		*t = i

		return err
	}

	return ErrBadTarget
}

func MapSliceFunc(src string, target interface{}) error {
	switch t := target.(type) {
	case *[]string:
		*t = strings.Split(src, ",")

		return nil
	}

	return ErrBadTarget
}

func MapStringFunc(src string, target interface{}) error {
	if t, ok := target.(*string); ok {
		*t = src

		return nil
	}

	return ErrBadTarget
}

func MapFloatFunc(src string, target interface{}) error {
	switch t := target.(type) {
	case *float32:
		f, err := strconv.ParseFloat(src, 32)
		*t = float32(f)

		return err
	case *float64:
		f, err := strconv.ParseFloat(src, 64)
		*t = f

		return err
	}

	return ErrBadTarget
}

func MapScannerFunc(src string, target interface{}) error {
	if t, ok := target.(sql.Scanner); ok {
		return t.Scan(src)
	}

	return ErrBadTarget
}
