package customers

import (
    "labix.org/v2/mgo/bson"
)

// TODO: define this as tags
var (
	searchable = []string{
        "name",
        "email",
        "company",
        "address.city",
        "address.street",
	}
)

type Query bson.M

func (q *Query) Id(param string) {
	if len(param) != 0 {
		(*q)["_id"] = bson.ObjectIdHex(param)
	}
}

func (q *Query) Search(param string) {
	if len(param) != 0 {
		o := []bson.M{}
		r := bson.RegEx{param, "i"}

		for _, value := range(searchable) {
			c := bson.M{}
			c[value] = &r
			o = append(o, c)
		}

		(*q)["$or"] = o
	}
}

func (q *Query) Filter(param string) {

}
