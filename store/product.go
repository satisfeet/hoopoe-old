package store

import (
	"io"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

type Product struct {
	*store
}

func NewProduct(db *mgo.Database) *Product {
	return &Product{
		store: &store{db},
	}
}

func (s *Product) CreateImage(m Model) (io.ReadWriteCloser, error) {
	iid := bson.NewObjectId()
	pid := utils.GetFieldValue(m, "Id")

	f, err := s.filesystem(m).Create("")
	if err != nil {
		return nil, err
	}
	f.SetId(iid)

	u := query{}
	u.Push("images", iid)

	if err := s.collection(m).UpdateId(pid, u); err != nil {
		f.Close()

		return nil, err
	}

	return f, nil
}

func (s *Product) OpenImage(m Model, iid bson.ObjectId) (io.ReadWriteCloser, error) {
	pid := utils.GetFieldValue(m, "Id")

	if err := s.collection(m).FindId(pid).One(nil); err != nil {
		return nil, err
	}

	return s.filesystem(m).OpenId(iid)
}

func (s *Product) RemoveImage(m Model, iid bson.ObjectId) error {
	pid := utils.GetFieldValue(m, "Id")

	u := query{}
	u.Pull("images", iid)

	if err := s.collection(m).UpdateId(pid, u); err != nil {
		return err
	}

	return s.filesystem(m).RemoveId(iid)
}
