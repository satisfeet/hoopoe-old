package mongo

import (
	"errors"
	"reflect"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const TagName = "store"

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
	ErrNotConnected   = errors.New("not connected")
	ErrStillConnected = errors.New("still connected")
)

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
	var n string

	switch v := value(model); v.Kind() {
	case reflect.Struct:
		n = v.Type().Name()
	case reflect.Array, reflect.Slice:
		n = v.Type().Elem().Name()
	}

	return store.database.C(strings.ToLower(n) + "s")
}

type fieldInfo struct {
	Index  bool
	Unique bool
}

type structInfo map[string]fieldInfo

func getStructInfo(m interface{}) structInfo {
	si := make(structInfo)

	switch v := reflect.ValueOf(m); v.Kind() {
	case reflect.Ptr, reflect.Array, reflect.Slice:
		v = v.Elem()

		fallthrough
	case reflect.Struct:
		t := v.Type()

		for n := 0; n < t.NumField(); n++ {
			f := t.Field(n)

			if t := f.Tag.Get(TagName); len(t) > 0 {
				si[strings.ToLower(f.Name)] = fieldInfo{
					Index:  strings.Contains(t, "index"),
					Unique: strings.Contains(t, "unique"),
				}
			}

			if f.Type.Kind() == reflect.Struct {
				nsi := getStructInfo(v.FieldByName(f.Name).Interface())

				for n, nsi := range nsi {
					si[strings.ToLower(f.Name+"."+n)] = nsi
				}
			}
		}
	}

	return si
}

func (store *Store) Index(model interface{}) error {
	s := store.clone()
	defer s.Close()

	i := mgo.Index{
		Key: make([]string, 0),
	}
	u := mgo.Index{
		Key:    make([]string, 0),
		Unique: true,
	}

	for k, v := range getStructInfo(model) {
		switch {
		case v.Index:
			i.Key = append(i.Key, k)
		case v.Unique:
			u.Key = append(u.Key, k)
		}
	}

	if len(i.Key) > 0 {
		if err := store.collection(model).With(s).EnsureIndex(i); err != nil {
			return err
		}
	}
	if len(u.Key) > 0 {
		return store.collection(model).With(s).EnsureIndex(u)
	}

	return nil
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

	switch v := value(model); v.Kind() {
	case reflect.Struct:
		f := v.FieldByName("Id")

		if id := f.Interface().(bson.ObjectId); !id.Valid() {
			f.Set(reflect.ValueOf(bson.NewObjectId()))
		}
	}

	return store.collection(model).With(s).Insert(model)
}

func (store *Store) Update(model interface{}) error {
	q := make(Query)
	s := store.clone()
	defer s.Close()

	switch v := value(model); v.Kind() {
	case reflect.Struct:
		q.Id(v.FieldByName("Id").Interface())
	}

	return store.collection(model).With(s).Update(q, model)
}

func (store *Store) Remove(model interface{}) error {
	s := store.clone()
	defer s.Close()

	return store.collection(model).With(s).Remove(model)
}

func value(m interface{}) reflect.Value {
	if v := reflect.ValueOf(m); v.Kind() == reflect.Ptr {
		return v.Elem()
	} else {
		return v
	}
}
