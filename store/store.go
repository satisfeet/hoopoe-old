package store

import (
	"reflect"
	"strings"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/mongo"
)

const TagName = "store"

type Store struct {
	mongo *mongo.Store
}

func NewStore() *Store {
	m := &mongo.Store{}

	return &Store{
		mongo: m,
	}
}

func (s *Store) Dial(u string) error {
	if err := s.mongo.Dial(u); err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() error {
	if err := s.mongo.Close(); err != nil {
		return err
	}
	return nil
}

func (store *Store) Search(pattern string, models interface{}) error {
	var t reflect.Type
	switch v := reflect.ValueOf(models); v.Kind() {
	case reflect.Ptr:
		v = v.Elem()
		fallthrough
	case reflect.Slice, reflect.Array:
		t = v.Elem().Type()
	}

	q := make(mongo.Query)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		if strings.Contains(f.Tag.Get(TagName), "search") {
			sq := make(mongo.Query)

			if err := sq.Regex(strings.ToLower(f.Name), pattern); err != nil {
				return err
			}

			q.Or(sq)
		}
	}

	return store.mongo.Find(q, models)
}

func (store *Store) FindAll(models interface{}) error {
	q := make(mongo.Query)

	return store.mongo.Find(q, models)
}

func (store *Store) FindId(id, model interface{}) error {
	q := make(mongo.Query)

	if err := q.Id(id); err != nil {
		return err
	}

	return store.mongo.FindOne(q, model)
}

func (store *Store) Insert(model interface{}) error {
	if err := validation.Validate(model); err != nil {
		return err
	}

	return store.mongo.Insert(model)
}

func (store *Store) Update(model interface{}) error {
	if err := validation.Validate(model); err != nil {
		return err
	}

	return store.mongo.Update(model)
}

func (store *Store) Remove(model interface{}) error {
	return store.mongo.Remove(model)
}
