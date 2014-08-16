package store

import (
	"io"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

type Order struct {
	*store
}

func NewOrder(s *mgo.Session) *Order {
	return &Order{
		store: &store{
			session:  s,
			database: s.DB(""),
		},
	}
}

func (s *Order) FindCustomer(o *model.Order) error {
	o.Customer.Id = o.CustomerId

	return s.FindOne(&o.Customer)
}

func (s *Order) FindProducts(o *model.Order) error {
	q := bson.M{}
	or := []bson.M{}

	for _, i := range o.Items {
		q := bson.M{}
		q["_id"] = i.ProductId

		or = append(or, q)
	}
	q["$or"] = or

	p := []model.Product{}

	if err := s.collection(p).Find(q).All(&p); err != nil {
		return err
	}
	for i, p := range p {
		o.Items[i].Product = p
	}

	return nil
}

func (s *Order) ReadInvoice(o *model.Order, w io.Writer) error {
	if !o.Id.Valid() {
		return ErrBadId
	}

	f, err := s.files(o).OpenId(o.Id)

	if err != nil {
		return err
	}

	if _, err := io.Copy(w, f); err != nil {
		return err
	}

	return f.Close()
}

func (s *Order) WriteInvoice(o *model.Order, r io.Reader) error {
	if !o.Id.Valid() {
		return ErrBadId
	}

	f, err := s.files(o).Create("")
	f.SetId(o.Id)

	if err != nil {
		return err
	}

	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	return f.Close()
}
