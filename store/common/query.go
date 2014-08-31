package common

import "strconv"

type query interface {
	Where(field, value string)

	Limit(int)

	String() string
	Params() []interface{}
}

type Query struct {
	table string
	limit int
	where map[string]interface{}
}

func NewQuery(table string) *Query {
	return &Query{
		table: table,
		where: make(map[string]interface{}),
	}
}

func (q *Query) Where(field, value string) {
	q.where[field] = value
}

func (q *Query) Limit(n int) {
	q.limit = n
}

func (q *Query) String() string {
	sql := "SELECT * FROM " + q.table

	if l := len(q.where); l > 0 {
		sql += " WHERE "

		for k, _ := range q.where {
			sql += k + "=?"

			if l--; l != 0 {
				sql += " AND "
			}
		}
	}

	if q.limit != 0 {
		sql += " LIMIT " + strconv.Itoa(q.limit)
	}

	return sql
}

func (q *Query) Params() []interface{} {
	p := make([]interface{}, 0)

	for _, v := range q.where {
		p = append(p, v)
	}

	return p
}
