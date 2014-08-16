package store

import "fmt"

// The Address type represents a postal address placed in Germany.
type Address struct {
	City    string `validate:"required,min=4,max=60"`
	Street  string `validate:"min=4,max=60"`
	ZipCode int    `validate:"min=10000,max=99999"`
}

// The Pricing type represents a pricing data with later support of tax and
// discount.
//
// We store prices as integers to avoid float pointer arthimetic.
type Pricing struct {
	Retail int64 `validate:"required,min=1"`
}

// Returns retail price with two decimal places.
func (p Pricing) Float() float64 {
	return float64(p.Retail) / 100
}

// Implements Stringer interface and returns retail price as euro
// representation.
//
// This is likely used in templates to directly be printed with decimal places
// and euro sign.
func (p Pricing) String() string {
	return fmt.Sprintf("%.2f €", p.Float())
}

// The Variation type represents attributes of a product which may differ like
// color and size. Together with an assigned product it is the smalled unit in
// the inventory.
//
// This type may be extended in the future to support inventary stats.
type Variation struct {
	Size  string `validate:"required,len=5"`
	Color string `validate:"required,min=3"`
}
