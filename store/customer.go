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

	stmt, err := tx.Prepare(`
		INSERT INTO city (name)
		VALUES (?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(m.Address.City)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(`
		INSERT INTO address (street, code, city_id)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`)

	if err != nil {
		return err
	}

	res, err = stmt.Exec(m.Address.Street, m.Address.Code, id)

	if err != nil {
		return err
	}

	id, err = res.LastInsertId()

	if err != nil {
		return err
	}

	stmt, err = tx.Prepare(`
		INSERT INTO customer (name, email, address_id)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
	`)

	if err != nil {
		return err
	}

	res, err = stmt.Exec(m.Name, m.Email, id)

	if err != nil {
		return err
	}

	m.Id, err = res.LastInsertId()

	if err != nil {
		return err
	}

	return tx.Commit()
}
