package mongo

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

type Query bson.M

var (
	ErrBadQueryParam = errors.New("bad query id")
	ErrBadQueryValue = errors.New("bad query value")
)

func (q Query) Id(v interface{}) error {
	switch t := v.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			q["_id"] = bson.ObjectIdHex(t)

			return nil
		}
	case bson.ObjectId:
		if t.Valid() {
			q["_id"] = t

			return nil
		}
	}

	return ErrBadQueryParam
}

func (q Query) Or(c Query) error {
	if q["$or"] == nil {
		q["$or"] = make([]Query, 0)
	}
	if or, ok := q["$or"].([]Query); ok {
		q["$or"] = append(or, c)

		return nil
	}

	return ErrBadQueryValue
}

func (q Query) Regex(k, v string) error {
	if len(v) == 0 {
		return ErrBadQueryParam
	}
	q[k] = bson.RegEx{v, "i"}

	return nil
}
