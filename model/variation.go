package model

import (
	"strings"

	"github.com/satisfeet/hoopoe/model/store"
)

type Variation struct {
	Size  string
	Color string
}

type Variations []Variation

func (v *Variations) Scan(src interface{}) error {
	switch t := src.(type) {
	case string:
		for _, s := range strings.Split(t, ",") {
			a := strings.Split(s, ":")

			*v = append(*v, Variation{
				Size:  a[1],
				Color: a[0],
			})
		}

		return nil
	}

	return store.ErrBadScanType
}
