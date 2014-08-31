package store

import (
	"encoding/json"

	"github.com/jmoiron/sqlx"
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          int64
	Title       string
	Subtitle    string
	Description string
	Pricing
	Categories Categories
	Variations Variations
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type ProductStore struct {
	session *Session
}

func NewProductStore(s *Session) *ProductStore {
	return &ProductStore{
		session: s,
	}
}

func (s *ProductStore) Find(m *[]Product) error {
	return s.sqlx().Select(m, `SELECT * FROM product_variation_category`)
}

func (s *ProductStore) sqlx() *sqlx.DB {
	return sqlx.NewDb(s.session.database, Driver)
}
