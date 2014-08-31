package store

import (
	"encoding/json"
	"fmt"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/utils"
)

type Customer struct {
	Id    int64
	Name  string  `validate:"required,min=5"`
	Email *string `validate:"required,email"`
	Address
}

func (c Customer) Validate() error {
	return validation.Validate(c)
}

func (c Customer) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(c))
}

type CustomerQuery struct {
	table string
	where map[string]string
}

func NewCustomerQuery() *CustomerQuery {
	return &CustomerQuery{
		table: `customer_address_city`,
		where: make(map[string]string),
	}
}

func (q *CustomerQuery) Where(field, value string) {
	q.where[field] = value
}

func (q *CustomerQuery) String() string {
	sql := fmt.Sprintf("SELECT * FROM %s", q.table)

	if l := len(q.where); l > 0 {
		sql += " WHERE "

		for k, v := range q.where {
			sql += fmt.Sprintf("%s = %s", k, v)

			if l--; l != 0 {
				sql += " AND "
			}
		}
	}

	fmt.Printf("SQL: %s\n", sql)

	return sql
}

type CustomerStore struct {
	session *common.Session
}

func NewCustomerStore(s *common.Session) *CustomerStore {
	return &CustomerStore{
		session: s,
	}
}

func (s *CustomerStore) Find(q *CustomerQuery, m *[]Customer) error {
	return s.session.Select(q.String(), m)
}

func (s *CustomerStore) FindOne(q *CustomerQuery, m *Customer) error {
	return s.session.SelectOne(q.String(), m)
}
