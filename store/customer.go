package store

import (
	"database/sql"
	"encoding/json"

	"github.com/satisfeet/go-validation"
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
	db *sql.DB
}

func NewCustomerStore(db *sql.DB) *CustomerStore {
	return &CustomerStore{
		db: db,
	}
}

func (s *CustomerStore) Find(m *[]Customer) error {
	rows, err := s.db.Query(`
		SELECT id, name, email, street, code, city
		FROM customer_address_city
	`)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		c := Customer{}

		var email, street, city sql.NullString
		var code sql.NullInt64

		if err := rows.Scan(&c.Id, &c.Name, &email, &street, &code, &city); err != nil {
			return err
		}

		c.Email = email.String
		c.Address.Street = street.String
		c.Address.City = city.String
		c.Address.Code = int(code.Int64)

		*m = append(*m, c)
	}

	return rows.Err()
}

func (s *CustomerStore) FindId(id interface{}, m *Customer) error {
	row := s.db.QueryRow(`
		SELECT id, name, email, street, code, city
		FROM customer_address_city
		WHERE id=?
	`, id)

	c := Customer{}

	var email, street, city sql.NullString
	var code sql.NullInt64

	err := row.Scan(&c.Id, &c.Name, &email, &street, &code, &city)

	c.Email = email.String
	c.Address.Street = street.String
	c.Address.City = city.String
	c.Address.Code = int(code.Int64)

	return err
}

func (s *CustomerStore) Search(query string, m *[]Customer) error {
	if len(query) == 0 {
		return s.Find(m)
	}

	query = "'%" + query + "%'"

	rows, err := s.db.Query(`
		SELECT id, name, email, street, code, city
		FROM customer_address_city
		WHERE name LIKE ?
		OR email LIKE ?
		OR city LIKE ?
		OR street LIKE ?
	`, query, query, query, query)

	if err != nil {
		return err
	}

	defer rows.Close()

	return scanToSlice(rows, m)
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
