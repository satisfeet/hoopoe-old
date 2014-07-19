package store

import "gopkg.in/mgo.v2/bson"

type Query bson.M

func (q Query) Id(s string) {
	if bson.IsObjectIdHex(s) {
		q["_id"] = bson.ObjectIdHex(s)
	}
}

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

func (q Query) Valid() bool {
	return q["_id"].(bson.ObjectId).Valid()
}