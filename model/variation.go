package model

import (
	"database/sql"
	"encoding/json"

	"strings"

	"github.com/satisfeet/hoopoe/utils"
)

type Variation struct {
	Size  string
	Color string
}

func (v Variation) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(v))
}

type Variations []Variation

func (v *Variations) Scan(src interface{}) error {
	switch t := src.(type) {
	case sql.RawBytes:
		for _, s := range strings.Split(string(t), ",") {
			a := strings.Split(s, ":")

			*v = append(*v, Variation{
				Size:  a[1],
				Color: a[0],
			})
		}

		return nil
	}

	return ErrBadScanType
}
