package store

import (
	"github.com/satisfeet/hoopoe/model"
	"github.com/satisfeet/hoopoe/store/mongo"
)

type Order struct {
	*store
}

func NewOrder(s *mongo.Store) *Order {
	return &Order{
		store: &store{s},
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
