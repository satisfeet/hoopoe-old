package model

import "fmt"

type Pricing struct {
	Retail int64 `json:"retail" validate:"required,min=1"`
}

func (p Pricing) Float() float64 {
	return float64(p.Retail) / 100
}

func (p Pricing) String() string {
	return fmt.Sprintf("%.2f €", p.Float())
}
