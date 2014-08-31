package store

import "strconv"

type Pricing struct {
	Retail float64
}

func (p *Pricing) Scan(src interface{}) error {
	switch t := src.(type) {
	case []byte:
		f, err := strconv.ParseFloat(string(t), 64)

		if err != nil {
			return err
		}

		p.Retail = f

		return nil
	}

	return ErrBadScanType
}
