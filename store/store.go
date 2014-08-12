package store

import (
	"errors"
	"strings"

	"github.com/satisfeet/hoopoe/utils"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var ErrBadParam = errors.New("bad param")

type Model interface{}

type store struct {
	Mongo *mgo.Database
}

func (s *store) Find(m Model) error {
	return s.collection(m).Find(nil).All(m)
}

func (s *store) FindId(id bson.ObjectId, m Model) error {
	return s.collection(m).FindId(id).One(m)
}

func (s *store) Insert(m Model) error {
	return s.collection(m).Insert(m)
}

func (s *store) Update(m Model) error {
	id := utils.GetFieldValue(m, "Id")

	return s.collection(m).UpdateId(id, m)
}

func (s *store) Remove(m Model) error {
	id := utils.GetFieldValue(m, "Id")

	return s.collection(m).RemoveId(id)
}

func (s *store) filesystem(m Model) *mgo.GridFS {
	n := strings.ToLower(utils.GetTypeName(m)) + "s"

	return s.Mongo.GridFS(n)
}

func (s *store) collection(m Model) *mgo.Collection {
	n := strings.ToLower(utils.GetTypeName(m)) + "s"

	return s.Mongo.C(n)
}

type query bson.M

func (q query) Id(id bson.ObjectId) {
	q["_id"] = id
}

func (q query) In(field string, value interface{}) {
	q[field] = bson.M{"$in": []interface{}{value}}
}

func (q query) Push(field string, value interface{}) {
	b := bson.M{}
	b[field] = value

	q["$push"] = b
}

func (q query) Pull(field string, value interface{}) {
	b := bson.M{}
	b[field] = value

	q["$pull"] = b
}

func ParseId(id interface{}) bson.ObjectId {
	var oid bson.ObjectId

	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			oid = bson.ObjectIdHex(t)
		}
	case bson.ObjectId:
		if t.Valid() {
			oid = t
		}
	}

	return oid
}
