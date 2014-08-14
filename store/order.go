package store

import (
	"io"

	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
	"gopkg.in/mgo.v2"
)

type Order struct {
	*store
}

func NewOrder(m *mongo.Store) *Order {
	s := &store{m}

	return &Order{
		store: s,
	}
}

func (s *Order) FindCustomer(o *model.Order) error {
	id := o.CustomerRef.Id

	return s.mongo.FindId(getName(o.Customer), id, &o.Customer)
}

func (s *Order) FindProducts(o *model.Order) error {
	q := mongo.Query{}
	or := []mongo.Query{}

	for _, i := range o.Items {
		q := mongo.Query{}
		q.Id(i.ProductRef.Id)

		or = append(or, q)
	}

	q["$or"] = or

	p := []model.Product{}

	if err := s.mongo.Find(getName(p), q, &p); err != nil {
		return err
	}

	for i, p := range p {
		o.Items[i].Product = p
	}

	return nil
}

func (s *Order) ReadInvoice(o *model.Order) (io.ReadCloser, error) {
	return s.mongo.OpenFileId(getName(o), o.Id)
}

func (s *Order) WriteInvoice(o *model.Order) (io.WriteCloser, error) {
	f, err := s.mongo.OpenFileId(getName(o), o.Id)

	if err == mgo.ErrNotFound {
		f, err = s.mongo.CreateFile(getName(o))
	}

	return f, err
}
