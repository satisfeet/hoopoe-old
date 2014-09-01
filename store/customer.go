package store

import (
	"database/sql"
	"encoding/json"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	Id      int64
	Name    string `validate:"required,min=5"`
	Email   string `validate:"required,email"`
	Address Address
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(c))
}

type CustomerStore struct {
	db    *sql.DB
	store *common.Store
}

func NewCustomerStore(db *sql.DB) *CustomerStore {
	return &CustomerStore{
		db:    db,
		store: common.NewStore(db),
	}
}

func (s *CustomerStore) Find(m *[]Customer) error {
	sql := `
		SELECT id, name, email, street, code, city
		FROM customer_address_city
	`

	return s.store.Query(sql).All(m)
}

func (s *CustomerStore) FindId(id interface{}, m *Customer) error {
	sql := `
		SELECT id, name, email, street, code, city
		FROM customer_address_city
		WHERE id=?
	`

	return s.store.Query(sql, id).One(m)
}

func (s *CustomerStore) Search(query string, m *[]Customer) error {
	if len(query) == 0 {
		return s.Find(m)
	}

	sql := `
		SELECT id, name, email, street, code, city
		FROM customer_address_city
		WHERE name LIKE ?
		OR email LIKE ?
		OR city LIKE ?
		OR street LIKE ?
	`

	p := make([]interface{}, 4)

	for i, _ := range p {
		p[i] = "'%" + query + "%'"
	}

	return s.store.Query(sql, p...).All(m)
}

func (s *CustomerStore) Insert(m *Customer) error {
	if err := validation.Validate(m); err != nil {
		return err
	}

	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	id, err := execPrepareId(tx, `
		INSERT INTO city (name)
		VALUES (?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`, m.Address.City)

	if err != nil {
		return err
	}

	id, err = execPrepareId(tx, `
		INSERT INTO address (street, code, city_id)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`, m.Address.Street, m.Address.Code, id)

	if err != nil {
		return err
	}

	id, err = execPrepareId(tx, `
		INSERT INTO customer (name, email, address_id)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`, m.Name, m.Email, id)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *CustomerStore) UpdateId(id interface{}, m *Customer) error {
	if err := validation.Validate(m); err != nil {
		return err
	}

	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	row := tx.QueryRow(`
		SELECT COUNT(id)
		FROM customer
		WHERE id = ?
	`, id)

	var n int64

	if err := row.Scan(&n); err != nil {
		tx.Rollback()

		return err
	}

	if n == 0 {
		tx.Rollback()

		return ErrNotFound
	}

	err = execPrepare(tx, `
		UPDATE customer
		SET name = ?, email = ?
		WHERE id = ?
	`, m.Name, m.Email, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	cid, err := execPrepareId(tx, `
		INSERT INTO city (name)
		VALUES (?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`, m.Address.City)

	if err != nil {
		tx.Rollback()

		return err
	}

	err = execPrepare(tx, `
		UPDATE address
		SET street = ?, code = ?, city_id = ?
		WHERE id=(SELECT address_id FROM customer WHERE id = ?)
	`, m.Address.Street, m.Address.Code, cid, id)

	if err != nil {
		tx.Rollback()

		return err
	}

	return tx.Commit()
}

func (s *CustomerStore) RemoveId(id interface{}) error {
	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	n, err := execPrepareAffected(tx, `
		DELETE FROM customer
		WHERE id = ?
	`, id)

	if err != nil {
		return err
	}

	if n == 0 {
		return ErrNotFound
	}

	return tx.Commit()
}
