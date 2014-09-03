package model

import (
	"database/sql"

	"github.com/satisfeet/hoopoe/model/store"
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

var sqlSelectProduct = `
	SELECT id, title, subtitle, description, price AS retail, variations, categories
	FROM product_variation_category
`

var sqlSelectProductId = `
	SELECT id, title, subtitle, description, price AS retail, variations, categories
	FROM product_variation_category
	WHERE id=?
`

type ProductStore struct {
	store *store.Store
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		store: store.NewStore(db),
	}
}

func (s *ProductStore) Find(m *[]Product) error {
	return s.store.Query(sqlSelectProduct).All(m)
}

func (s *ProductStore) FindId(id interface{}, m *Product) error {
	return s.store.Query(sqlSelectProductId, id).One(m)
}
