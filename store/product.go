package store

import (
	"encoding/json"
	"io"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

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

func NewProductStore(db *mgo.Database) *ProductStore {
	return &ProductStore{
		store: &store{db},
	}
}

func (s *ProductStore) CreateImage(m Product) (io.ReadWriteCloser, error) {
	image := bson.NewObjectId()

	f, err := s.filesystem(m).Create("")
	if err != nil {
		return nil, err
	}
	f.SetId(image)

	u := query{}
	u.Push("images", image)

	if err := s.collection(m).UpdateId(m.Id, u); err != nil {
		f.Close()

		return nil, err
	}

	return f, nil
}

func (s *ProductStore) OpenImage(m Product, image bson.ObjectId) (io.ReadWriteCloser, error) {
	if err := s.collection(m).FindId(m.Id).One(nil); err != nil {
		return nil, err
	}

	return s.filesystem(m).OpenId(image)
}

func (s *ProductStore) RemoveImage(m Product, image bson.ObjectId) error {
	u := query{}
	u.Pull("images", image)

	if err := s.collection(m).UpdateId(m.Id, u); err != nil {
		return err
	}

	return s.filesystem(m).RemoveId(image)
}
