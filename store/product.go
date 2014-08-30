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
	*Pricing
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
	return s.sqlx().Select(m, `
		SELECT pr.id, pr.title, pr.subtitle, pr.description, pr.price
		FROM product AS pr
	`)
}

func (s *ProductStore) sqlx() *sqlx.DB {
	return sqlx.NewDb(s.session.database, Driver)
}
