package mongo

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// The Query type helps you build mongo queries with an easy to use interface.
type Query struct {
	err    error
	query  bson.M
	limit  int
	offset int
}

// Error returned on invalid object id.
var ErrBadId = errors.New("bad id")

// Returns an new initialized Query.
func NewQuery() *Query {
	return &Query{
		query: make(bson.M),
	}
}

// Applies an equals id condition if id is valid object id. Does cast from
// string to object id when possible if not it sets an error.
func (q *Query) Id(id interface{}) {
	var oid bson.ObjectId

	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			oid = bson.ObjectIdHex(t)
		}
	case bson.ObjectId:
		oid = t
	}

	if oid.Valid() {
		q.query["_id"] = oid
	} else {
		q.err = ErrBadId
	}
}

// Applies an other query as optional condition.
func (q *Query) Or(query *Query) {
	or, ok := q.query["$or"].([]bson.M)

	if !ok {
		or = make([]bson.M, 0)
	}

	q.query["$or"] = append(or, query.query)
}

// Applies an matches regular expression condition.
func (q *Query) RegEx(field string, value string) {
	q.query[field] = bson.RegEx{value, "i"}
}

// Applies an equals value condition.
func (q *Query) Equals(field string, value interface{}) {
	q.query[field] = value
}

// Returns last error which occured or nil.
func (q *Query) Err() error {
	return q.err
}

// Defines limit for result sets to be fetched.
func (q *Query) SetLimit(n int) {
	q.limit = n
}

// Defines offset for result sets to be fetched.
func (q *Query) SetOffset(n int) {
	q.offset = n
}
