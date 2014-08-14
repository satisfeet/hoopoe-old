package store

import (
	"io"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Product struct {
	*store
}

func NewProduct(m *mongo.Store) *Product {
	s := &store{m}

	return &Product{
		store: s,
	}
}

func (s *Product) CreateImage(m *model.Product) (io.ReadWriteCloser, error) {
	id := bson.NewObjectId()

	f, err := s.mongo.CreateFile(getName(m))
	if err != nil {
		return nil, err
	}

	u := mongo.Query{}
	u.Push("images", id)

	if err := s.mongo.UpdateId(getName(m), m.Id, u); err != nil {
		f.Close()

		return nil, err
	}

	return f, nil
}

func (s *Product) OpenImage(m *model.Product, id interface{}) (io.ReadWriteCloser, error) {
	if err := s.FindOne(m); err != nil {
		return nil, err
	}

	return s.mongo.OpenFileId(getName(m), id)
}

func (s *Product) RemoveImage(m *model.Product, id interface{}) error {
	u := mongo.Query{}
	u.Pull("images", id)

	if err := s.mongo.UpdateId(getName(m), m.Id, u); err != nil {
		return err
	}

	return s.mongo.RemoveFileId(getName(m), id)
}
