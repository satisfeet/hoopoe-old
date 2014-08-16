package common

import "github.com/satisfeet/hoopoe/store/mongo"

// The Query type is the basic implementation of the query interface and is
// intended to be used as base Query for composition.
type Query struct {
	m *mongo.Query
}

// Returns initialized Query.
func NewQuery() *Query {
	return &Query{
		m: mongo.NewQuery(),
	}
}

// Applies an equals id query to all supported engines.
func (q *Query) Id(id interface{}) {
	q.mongo().Id(id)
}

// Applies an or condition to all supported engines.
func (q *Query) Or(or query) {
	q.mongo().Or(or.mongo())
}

// Applies an regex condition to all supported engines.
func (q *Query) RegEx(field, value string) {
	q.mongo().RegEx(field, value)
}

// Returns the query compatible with the mongodb engine.
func (q *Query) mongo() *mongo.Query {
	return q.m
}
