package model

type Address struct {
	City    string `json:"city"    "omitempty"`
	Street  string `json:"street"  "omitempty"`
	Zipcode int    `json:"zipcode" "omitempty"`
}
