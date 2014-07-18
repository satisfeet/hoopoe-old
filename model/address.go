package model

type Address struct {
	Zip    int    `json:"zip,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
}
