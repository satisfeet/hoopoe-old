package store

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Query extends a bson map to do common mongodb queries.
type Query bson.M

var (
	ErrBadIdQuery     = errors.New("bad id query")
	ErrBadSearchQuery = errors.New("bad search query")
)

// Id takes a string which will be used as object id condition if valid.
func (q Query) Id(s string) error {
	if bson.IsObjectIdHex(s) {
		q["_id"] = bson.ObjectIdHex(s)
	} else {
		return ErrBadIdQuery
	}
	return nil
}

// Search takes a string which will be used as $regex in an $or condition
// accross all defined field names.
func (q Query) Search(s string, f []string) error {
	if len(s) == 0 || len(f) == 0 {
		return ErrBadSearchQuery
	}

	r := bson.RegEx{s, "i"}
	o := make([]bson.M, len(f))

	for i := 0; i < len(f); i++ {
		o[i] = make(bson.M, 1)
		o[i][f[i]] = r
	}
	q["$or"] = o

	return nil
}
