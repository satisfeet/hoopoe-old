package model

type Pricing struct {
	Retail int64 `json:"retail" validate:"required,min=1"`
}
