package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/satisfeet/hoopoe/utils"
)

var ErrBadVariationType = errors.New("bad variation type")

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
	case []byte:
		for _, s := range strings.Split(string(t), ",") {
			a := strings.Split(s, ":")

			*v = append(*v, Variation{
				Size:  a[1],
				Color: a[0],
			})
		}

		return nil
	default:
		fmt.Printf("variation: %v\n", t)
	}

	return ErrBadVariationType
}
