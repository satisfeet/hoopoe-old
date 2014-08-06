package model

import "encoding/json"

type Address struct {
	City    string `validate:"required,min=4,max=60" store:"index"`
	Street  string `validate:"min=4,max=60" store:"index"`
	ZipCode int    `validate:"min=10000,max=99999"`
}

func (a Address) Marshal() map[string]interface{} {
	m := make(map[string]interface{})

	if len(a.City) != 0 {
		m["city"] = a.City
	}
	if len(a.Street) != 0 {
		m["street"] = a.Street
	}
	if a.ZipCode != 0 {
		m["zipcode"] = a.ZipCode
	}

	return m
}

func (a Address) MarshalJSON() ([]byte, error) {
	m := a.Marshal()

	return json.Marshal(m)
}
