package store

import (
    "labix.org/v2/mgo/bson"
)

type Query map[string]string

func (q *Query) Has(key string) bool {
    return len(q.Get(key)) != 0
}

func (q *Query) Get(key string) string {
    return (*q)[key]
}

func (q *Query) Bson(config map[string][]string) *bson.M {
    m := bson.M{}

    if q.Has("id") {
        m["_id"] = bson.ObjectIdHex(q.Get("id"))
    }

    if q.Has("search") {
        o := []bson.M{}
        r := bson.RegEx{q.Get("search"), "i"}

        for _, value := range(config["searchable"]) {
            c := bson.M{}

            c[value] = &r

            o = append(o, c)
        }

        m["$or"] = o
    }

    return &m
}
