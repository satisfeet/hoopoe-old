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

var sqlSelectCustomer = `
	SELECT id, name, email, street, code, city
	FROM customer_address_city
`

var sqlSelectCustomerId = `
	SELECT id, name, email, street, code, city
	FROM customer_address_city
	WHERE id=?
`

var sqlSelectCustomerCount = `
	SELECT COUNT(id)
	FROM customer
	WHERE id = ?
`

var sqlSelectCustomerSearch = `
	SELECT id, name, email, street, code, city
	FROM customer_address_city
	WHERE name LIKE ?
	OR email LIKE ?
	OR city LIKE ?
	OR street LIKE ?
`

var sqlInsertCustomer = `
	INSERT INTO customer (name, email, address_id)
	VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE id=LAST_INSERT_ID(id)
`

var sqlUpdateCustomer = `
	UPDATE customer
	SET name = ?, email = ?
	WHERE id = ?
`

var sqlDeleteCustomer = `
	DELETE FROM customer
	WHERE id = ?
`

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
	return s.store.Query(sqlSelectCustomer).All(m)
}

func (s *CustomerStore) FindId(id interface{}, m *Customer) error {
	return s.store.Query(sqlSelectCustomerId, id).One(m)
}

func (s *CustomerStore) Search(query string, m *[]Customer) error {
	if len(query) == 0 {
		return s.Find(m)
	}

	p := make([]interface{}, 4)

	for i, _ := range p {
		p[i] = "'%" + query + "%'"
	}

	return s.store.Query(sqlSelectCustomerSearch, p...).All(m)
}

func (s *CustomerStore) Insert(m *Customer) error {
	if err := validation.Validate(m); err != nil {
		return err
	}

	tx := s.store.Begin()

	res, err := tx.Exec(sqlInsertCity, m.Address.City)
	if err != nil {
		return err
	}

	res, err = tx.Exec(sqlInsertAddress, m.Address.Street, m.Address.Code, res.Id)
	if err != nil {
		return err
	}

	res, err = tx.Exec(sqlInsertCustomer, m.Name, m.Email, res.Id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *CustomerStore) UpdateId(id interface{}, m *Customer) error {
	if err := validation.Validate(m); err != nil {
		return err
	}

	tx := s.store.Begin()

	var n int64
	if err := tx.Tx.QueryRow(sqlSelectCustomerCount, id).Scan(&n); err != nil {
		tx.Rollback()

		return err
	}
	if n == 0 {
		tx.Rollback()

		return ErrNotFound
	}

	_, err := tx.Exec(sqlUpdateCustomer, m.Name, m.Email, id)
	if err != nil {
		return err
	}

	res, err := tx.Exec(sqlInsertCity, m.Address.City)
	if err != nil {
		return err
	}

	res, err = tx.Exec(sqlUpdateAddress, m.Address.Street, m.Address.Code, res.Id, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *CustomerStore) RemoveId(id interface{}) error {
	tx := s.store.Begin()

	res, err := tx.Exec(sqlDeleteCustomer, id)
	if err != nil {
		return err
	}
	if res.Rows == 0 {
		return ErrNotFound
	}

	return tx.Commit()
}
