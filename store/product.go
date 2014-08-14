package store

import (
	"io"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

type Product struct {
	*store
}

func NewProduct(s *mgo.Session) *Product {
	return &Product{
		store: &store{
			session:  s,
			database: s.DB(""),
		},
	}
}

func (s *Product) pushImage(p *model.Product, id bson.ObjectId) error {
	u := bson.M{"$pull": bson.M{"images": id}}

	if !p.Id.Valid() {
		return ErrBadId
	}

	c := s.session.Clone()
	defer c.Close()

	return s.collection(p).With(c).UpdateId(p.Id, u)
}

func (s *Product) pullImage(p *model.Product, id bson.ObjectId) error {
	u := bson.M{"$pull": bson.M{"images": id}}

	if !p.Id.Valid() {
		return ErrBadId
	}

	c := s.session.Clone()
	defer c.Close()

	return s.collection(p).With(c).UpdateId(p.Id, u)
}

func (s *Product) ReadImage(p *model.Product, id bson.ObjectId, w io.Writer) error {
	if !id.Valid() {
		return ErrBadId
	}
	if !p.Id.Valid() {
		return ErrBadId
	}

	if err := s.collection(p).FindId(p.Id).One(nil); err != nil {
		return err
	}

	f, err := s.files(p).OpenId(id)

	if err != nil {
		return err
	}

	if _, err := io.Copy(w, f); err != nil {
		return err
	}

	return f.Close()
}

func (s *Product) WriteImage(p *model.Product, id bson.ObjectId, r io.Reader) error {
	if !id.Valid() {
		return ErrBadId
	}

	f, err := s.files(p).Create("")
	f.SetId(id)

	if err != nil {
		return err
	}

	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	return s.pushImage(p, id)
}

func (s *Product) RemoveImage(p *model.Product, id bson.ObjectId) error {
	if !id.Valid() {
		return ErrBadId
	}

	if err := s.pullImage(p, id); err != nil {
		return err
	}

	return s.files(p).RemoveId(id)
}
