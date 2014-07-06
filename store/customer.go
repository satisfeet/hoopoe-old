package store

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	CustomerUnique = []string{
		"email",
	}
	CustomerIndices = []string{
		"name",
		"company",
		"address.street",
		"address.city",
	}
)

type Customer struct {
	Id      bson.ObjectId   `json:"id"     bson:"_id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Company string          `json:"company,omitempty"`
	Address CustomerAddress `json:"address"`
}

type CustomerAddress struct {
	Zip    uint16 `json:"zip,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
}

type CustomerHandler struct {
	store *Store
}

func NewCustomerHandler(s *Store) *CustomerHandler {
	return &CustomerHandler{s}
}

func (h *CustomerHandler) Index() {
	s := h.store.Mongo().Clone()

	defer s.Close()

	c := s.DB("").C("customers")

	c.EnsureIndex(mgo.Index{Key: CustomerIndices})
	c.EnsureIndex(mgo.Index{Key: CustomerUnique, Unique: true})
}

func (h *CustomerHandler) Create(c *Customer) error {
	s := h.store.Mongo().Clone()

	defer s.Close()

	if !c.Id.Valid() {
		c.Id = bson.NewObjectId()
	}

	return s.DB("").C("customers").Insert(c)
}

func (h *CustomerHandler) Update(c *Customer) error {
	s := h.store.Mongo().Clone()

	defer s.Close()

	return s.DB("").C("customers").UpdateId(c.Id, c)
}

func (h *CustomerHandler) Remove(q Query) error {
	s := h.store.Mongo().Clone()

	defer s.Close()

	return s.DB("").C("customers").Remove(q)
}

func (h *CustomerHandler) FindAll(q Query, c *[]Customer) error {
	s := h.store.Mongo().Clone()

	defer s.Close()

	return s.DB("").C("customers").Find(q).All(c)
}

func (h *CustomerHandler) FindOne(q Query, c *Customer) error {
	s := h.store.Mongo().Clone()

	defer s.Close()

	return s.DB("").C("customers").Find(q).One(c)
}
