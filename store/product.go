package store

import (
	"encoding/json"

	"github.com/satisfeet/go-validation"

	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

var productName = "products"

var productUnique = []string{
	"name",
}

type Product struct {
	Id          interface{}   `bson:"_id"`
	Name        string        `validate:"required,min=10,max=28"`
	Images      []interface{} `validate:"min=1"`
	Pricing     Pricing       `validate:"required"`
	Variations  []Variation   `validate:"required"`
	Description string        `validate:"required,min=40"`
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

func NewProductQuery() *common.Query {
	return common.NewQuery()
}

type ProductStore struct {
	*common.Store
}

func NewProductStore(s *common.Session) (*ProductStore, error) {
	ps := &ProductStore{
		Store: common.NewStore(common.Config{
			Name:   productName,
			Unique: productUnique,
		}, s),
	}

	return ps, ps.Index()
}
