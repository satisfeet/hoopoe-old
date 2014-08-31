package store

import (
	"database/sql"
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          int64
	Title       string
	Subtitle    string
	Description string
	Price       Pricing
	Categories  Categories
	Variations  Variations
}

func (p Product) Validate() error {
	return validation.Validate(p)
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type ProductStore struct {
	db *sql.DB
}

func NewProductStore(db *sql.DB) *ProductStore {
	return &ProductStore{
		db: db,
	}
}

func (s *ProductStore) Find(m *[]Product) error {
	rows, err := s.db.Query(`
		SELECT *
		FROM product_variation_category
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	return scanToSlice(rows, m)
}

func (s *ProductStore) FindId(id interface{}, m *Product) error {
	rows, err := s.db.Query(`
		SELECT *
		FROM product_variation_category
		WHERE id=?
	`, id)

	if err != nil {
		return err
	}

	defer rows.Close()

	return scanToSlice(rows, m)
}
