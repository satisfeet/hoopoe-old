package store

import (
	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/common"
	"github.com/satisfeet/hoopoe/store/mongo"
)

const TagName = "store"

type Store struct {
	mongo *mongo.Store
}

func NewStore() *Store {
	return &Store{
		mongo: &mongo.Store{},
	}
}

func (s *Store) Dial(url string) error {
	if err := s.mongo.Dial(url); err != nil {
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

func (s *Store) Search(pattern string, models interface{}) error {
	q := mongo.Query{}

	for k, _ := range common.GetStructInfo(models) {
		sq := mongo.Query{}

		if err := sq.Regex(k, pattern); err != nil {
			return err
		}

		q.Or(sq)
	}

	return s.mongo.Find(q, models)
}

func (s *Store) FindAll(models interface{}) error {
	q := make(mongo.Query)

	return s.mongo.Find(q, models)
}

func (s *Store) FindId(id, model interface{}) error {
	q := make(mongo.Query)

	if err := q.Id(id); err != nil {
		return err
	}

	return s.mongo.FindOne(q, model)
}

func (s *Store) Insert(model interface{}) error {
	if err := validation.Validate(model); err != nil {
		return err
	}

	return s.mongo.Insert(model)
}

func (s *Store) Update(model interface{}) error {
	if err := validation.Validate(model); err != nil {
		return err
	}

	return s.mongo.Update(model)
}

func (s *Store) Remove(model interface{}) error {
	return s.mongo.Remove(model)
}
