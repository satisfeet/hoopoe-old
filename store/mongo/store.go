package mongo

import (
	"errors"
	"reflect"
	"strings"

	"gopkg.in/mgo.v2"
)

const TagName = "store"

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

	switch t := value(model).Type(); t.Kind() {
	case reflect.Array, reflect.Slice:
		t = t.Elem()

		fallthrough
	case reflect.Struct:
		for n := 0; n < t.NumField(); n++ {
			f := t.Field(n)

			switch n := strings.ToLower(f.Name); f.Tag.Get(TagName) {
			case "-":
				i.Key = append(i.Key, n)
			case "unique":
				u.Key = append(u.Key, n)
			}
		}
	}

	if err := store.collection(model).With(s).EnsureIndex(i); err != nil {
		return err
	} else {
		return store.collection(model).With(s).EnsureIndex(u)
	}
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
