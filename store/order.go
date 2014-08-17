package store

import (
	"io"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model"
)

type Order struct {
	*store
	product  *Product
	customer *Customer
}

var OrderUnique = []string{
	"number",
}

var OrderName = "orders"

func NewOrder(s *mgo.Session) *Order {
	info := storeInfo{
		Name: OrderName,
	}

	return &Order{
		store: &store{
			info:     info,
			session:  s,
			database: s.DB(""),
		},
		product:  NewProduct(s),
		customer: NewCustomer(s),
	}
}

func (s *Order) Insert(o *model.Order) error {
	c := s.session.Clone()
	defer c.Close()

	if !o.Id.Valid() {
		o.Id = bson.NewObjectId()
	}

	if o.State.Created.IsZero() {
		o.State.Created = time.Now()
	}

	for o.Number = 1; o.Number != 0; o.Number++ {
		if err := o.Validate(); err != nil {
			return err
		}

		if err := s.store.collection().With(c).Insert(o); err != nil {
			if !mgo.IsDup(err) {
				return err
			}
		} else {
			break
		}
	}

	return nil
}

func (s *Order) FindCustomer(o *model.Order) error {
	o.Customer.Id = o.CustomerId

	return s.customer.FindOne(&o.Customer)
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

	if err := s.product.collection().Find(q).All(&p); err != nil {
		return err
	}
	for i, p := range p {
		o.Items[i].Product = p
	}

	return nil
}

func (s *Order) OpenInvoice(o *model.Order) (io.ReadCloser, error) {
	if !o.Id.Valid() {
		return nil, ErrBadId
	}

	return s.files().OpenId(o.Id)
}

func (s *Order) CreateInvoice(o *model.Order) (io.WriteCloser, error) {
	if !o.Id.Valid() {
		return nil, ErrBadId
	}

	f, err := s.files().Create("")
	f.SetId(o.Id)

	return f, err
}
