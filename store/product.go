package store

import (
	"database/sql"
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          int64
	Title       string
	Subtitle    string
	Description string
	Categories  []string
	Variations  Variations
	Pricing     Pricing
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type ProductStore struct {
	db    *sql.DB
	store *common.Store
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		db:    db,
		store: common.NewStore(db),
	}
}

func (s *ProductStore) Find(m *[]Product) error {
	sql := `
		SELECT id, title, subtitle, description, price AS retail, variations, categories
		FROM product_variation_category
	`

	return s.store.Query(sql).All(m)
}

func (s *ProductStore) FindId(id interface{}, m *Product) error {
	sql := `
		SELECT id, title, subtitle, description, price AS retail, variations, categories
		FROM product_variation_category
		WHERE id=?
	`

	return s.store.Query(sql, id).One(m)
}
