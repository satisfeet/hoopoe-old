package store

import (
	"errors"

	"github.com/satisfeet/hoopoe/store/common"
)

var ErrBadScanType = errors.New("bad scan type")

func Open(url string) (*common.Session, error) {
	s := &common.Session{}

	return s, s.Dial(url)
}
