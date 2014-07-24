package store

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

// Query extends a bson map to do common mongodb queries.
//
// NOTE: The main idea is that you can call methods on it which will then
// transform the given data into nested json structs.
type Query bson.M

// ErrInvalidQuery will be returned when a condition does not match the
// requirements. One common example is an executed equals-id-condition where
// we had an invalid id.
var ErrInvalidQuery = errors.New("invalid query")

// Id takes a string which will be used to setup an equals object id condition.
// If the string is no valid object id it will be ignored to not cause any
// errors in mongo or mgo.
func (q Query) Id(s string) {
	if bson.IsObjectIdHex(s) {
		q["_id"] = bson.ObjectIdHex(s)
	}
}

// Valid returns true if the query contains a valid id condition.
//
// TODO: Do not limit this onto the id parameter.
func (q Query) Valid() bool {
	if q["_id"] == nil {
		return false
	}

	return q["_id"].(bson.ObjectId).Valid()
}

// Search sets up a search string and an array of field names which will then
// used together in a regex-or-condition to emit a full-text search accross
// fields on the index.
//
// NOTE: As this is not very efficient the use should be for internal use only.
func (q Query) Search(s string, f []string) {
	if len(s) != 0 {
		r := bson.RegEx{s, "i"}
		o := make([]bson.M, len(f))

		for i := 0; i < len(f); i++ {
			o[i] = make(bson.M, 1)
			o[i][f[i]] = r
		}
		q["$or"] = o
	}
}
