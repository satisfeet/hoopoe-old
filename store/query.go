package store

import "gopkg.in/mgo.v2/bson"

// Query extends a bson map to do common mongodb queries.
type Query bson.M

// Id takes a string which will be used as object id condition if valid.
func (q Query) Id(s string) {
	if bson.IsObjectIdHex(s) {
		q["_id"] = bson.ObjectIdHex(s)
	}
}

// Valid returns true if the query contains a valid id condition.
//
// TODO: Valid only for ID makes not much sense. Rethink this.
func (q Query) Valid() bool {
	if q["_id"] == nil {
		return false
	}

	return q["_id"].(bson.ObjectId).Valid()
}

// Search takes a string which will be used as $regex in an $or condition
// accross all defined field names.
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
