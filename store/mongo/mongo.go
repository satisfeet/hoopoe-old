package mongo

import (
	"errors"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

type Query bson.M

var (
	ErrBadQueryParam = errors.New("bad query id")
	ErrBadQueryValue = errors.New("bad query value")
)

func (query Query) Id(id interface{}) error {
	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			query["_id"] = bson.ObjectIdHex(t)

			return nil
		}
	case bson.ObjectId:
		if t.Valid() {
			query["_id"] = t

			return nil
		}
	}

	return ErrBadQueryParam
}

func (query Query) Or(q Query) error {
	if query["$or"] == nil {
		query["$or"] = make([]Query, 0)
	}
	if or, ok := query["$or"].([]Query); ok {
		query["$or"] = append(or, q)

		return nil
	}

	return ErrBadQueryValue
}

func (query Query) Regex(key, pattern string) error {
	if len(key) == 0 || len(pattern) == 0 {
		return ErrBadQueryParam
	}
	query[key] = bson.RegEx{pattern, "i"}

	return nil
}

type Store struct {
	session  *mgo.Session
	database *mgo.Database
}

var (
	ErrBadModel       = errors.New("bad model")
	ErrNotConnected   = errors.New("not connected")
	ErrStillConnected = errors.New("still connected")
)

const IdFieldName = "Id"

func (store *Store) Dial(url string) error {
	if store.session != nil {
		return ErrStillConnected
	}

	s, err := mgo.Dial(url)

	if err != nil {
		return err
	}

	store.session = s
	store.database = s.DB("")

	return nil
}

func (store *Store) Close() error {
	if store.session == nil {
		return ErrNotConnected
	}

	store.session.Close()
	store.session = nil
	store.database = nil

	return nil
}

func (store *Store) clone() *mgo.Session {
	return store.session.Clone()
}

func (store *Store) collection(model interface{}) *mgo.Collection {
	n := utils.GetTypeName(model)

	return store.database.C(strings.ToLower(n) + "s")
}

func (store *Store) Index(model interface{}) error {
	s := store.clone()
	defer s.Close()

	i := make([]string, 0)
	u := make([]string, 0)

	for k, v := range utils.GetStructInfo(model) {
		switch {
		case v.Index:
			i = append(i, k)
		case v.Unique:
			u = append(u, k)
		}
	}

	var err error

	if len(u) > 0 {
		err = store.collection(model).With(s).EnsureIndex(mgo.Index{
			Key:    u,
			Unique: true,
		})
	}
	if err == nil && len(i) > 0 {
		err = store.collection(model).With(s).EnsureIndexKey(i...)
	}

	return err
}

func (store *Store) Find(query Query, model interface{}) error {
	s := store.clone()
	defer s.Close()

	return store.collection(model).With(s).Find(query).All(model)
}

func (store *Store) FindOne(query Query, model interface{}) error {
	s := store.clone()
	defer s.Close()

	return store.collection(model).With(s).Find(query).One(model)
}

func (store *Store) Insert(model interface{}) error {
	s := store.clone()
	defer s.Close()

	if id, ok := utils.GetFieldValue(model, IdFieldName).(bson.ObjectId); ok {
		if !id.Valid() {
			utils.SetFieldValue(model, IdFieldName, bson.NewObjectId())
		}
	} else {
		return ErrBadModel
	}

	return store.collection(model).With(s).Insert(model)
}

func (store *Store) Update(model interface{}) error {
	s := store.clone()
	defer s.Close()

	q := Query{}
	q.Id(utils.GetFieldValue(model, "Id"))

	return store.collection(model).With(s).Update(q, model)
}

func (store *Store) Remove(model interface{}) error {
	s := store.clone()
	defer s.Close()

	return store.collection(model).With(s).Remove(model)
}
