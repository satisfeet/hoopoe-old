package store

import (
	"io"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/mongo"
)

type Product struct {
	*store
}

func NewProduct(s *mongo.Store) *Product {
	return &Product{
		store: &store{s},
	}
}

func (s *Product) RemoveId(id interface{}) error {
	return s.mongo.RemoveId("products", id)
}

func (s *Product) CreateImage(pid interface{}) (io.ReadWriteCloser, error) {
	id := bson.NewObjectId()

	f, err := s.mongo.CreateFile("products")
	if err != nil {
		return nil, err
	}

	u := mongo.Query{}
	u.Push("images", id)

	if err := s.mongo.UpdateId("products", pid, u); err != nil {
		f.Close()

		return nil, err
	}

	return f, nil
}

func (s *Product) OpenImage(pid interface{}, iid interface{}) (io.ReadWriteCloser, error) {
	if err := s.mongo.FindId("products", pid, nil); err != nil {
		return nil, err
	}

	return s.mongo.OpenFileId("products", iid)
}

func (s *Product) RemoveImage(pid interface{}, iid interface{}) error {
	u := mongo.Query{}
	u.Pull("images", iid)

	if err := s.mongo.UpdateId("products", pid, u); err != nil {
		return err
	}

	return s.mongo.RemoveFileId("products", iid)
}
