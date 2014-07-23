package model

type Pricing struct {
	Retail int64 `json:"retail" validate:"nonzero,min=1"`
}
