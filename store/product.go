package store

import (
	"encoding/json"
	"io"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/mongo"
	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	Id          bson.ObjectId   `bson:"_id"`
	Name        string          `validate:"required,min=10,max=20"`
	Images      []bson.ObjectId `validate:"min=1"`
	Pricing     Pricing         `validate:"required,nested"`
	Variations  []Variation     `validate:"required,nested"`
	Description string          `validate:"required,min=40"`
}

func (p Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(utils.GetFieldValues(p))
}

type Variation struct {
	Size  string `validate:"required,len=5"`
	Color string `validate:"required,min=3"`
}

type ProductStore struct {
	*store
}

func NewProductStore(s *mongo.Store) *ProductStore {
	return &ProductStore{
		store: &store{s},
	}
}

func (s *ProductStore) RemoveId(id interface{}) error {
	return s.mongo.RemoveId("products", id)
}

func (s *ProductStore) CreateImage(pid interface{}) (io.ReadWriteCloser, error) {
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

func (s *ProductStore) OpenImage(pid interface{}, iid interface{}) (io.ReadWriteCloser, error) {
	if err := s.mongo.FindId("products", pid, nil); err != nil {
		return nil, err
	}

	return s.mongo.OpenFileId("products", iid)
}

func (s *ProductStore) RemoveImage(pid interface{}, iid interface{}) error {
	u := mongo.Query{}
	u.Pull("images", iid)

	if err := s.mongo.UpdateId("products", pid, u); err != nil {
		return err
	}

	return s.mongo.RemoveFileId("products", iid)
}
