package store

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Store struct {
	mongo *mgo.Session
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Open(url string) error {
	var err error

	s.mongo, err = mgo.Dial(url)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) IndexProduct() {
	c := s.mongo.DB(ProductDatabase).C(ProductCollection)

	c.EnsureIndex(mgo.Index{Key: ProductUnique, Unique: true})
}

func (s *Store) InsertProduct(p *Product) error {
	m := s.mongo.Clone()
	defer m.Close()

	if !p.Id.Valid() {
		p.Id = bson.NewObjectId()
	}

	return m.DB(ProductDatabase).C(ProductCollection).Insert(p)
}

func (s *Store) UpdateProduct(p *Product) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(ProductDatabase).C(ProductCollection).UpdateId(p.Id, p)
}

func (s *Store) RemoveProduct(q Query) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(ProductDatabase).C(ProductCollection).Remove(q)
}

func (s *Store) FindAllProduct(q Query, p *[]Product) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(ProductDatabase).C(ProductCollection).Find(q).All(p)
}

func (s *Store) FindOneProduct(q Query, p *Product) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(ProductDatabase).C(ProductCollection).Find(q).One(p)
}

func (s *Store) IndexCustomer() {
	c := s.mongo.DB(CustomerDatabase).C(CustomerCollection)

	c.EnsureIndex(mgo.Index{Key: CustomerIndices})
	c.EnsureIndex(mgo.Index{Key: CustomerUnique, Unique: true})
}

func (s *Store) InsertCustomer(c *Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	if !c.Id.Valid() {
		c.Id = bson.NewObjectId()
	}

	return m.DB(CustomerDatabase).C(CustomerCollection).Insert(c)
}

func (s *Store) UpdateCustomer(c *Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(CustomerDatabase).C(CustomerCollection).UpdateId(c.Id, c)
}

func (s *Store) RemoveCustomer(q Query) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(CustomerDatabase).C(CustomerCollection).Remove(q)
}

func (s *Store) FindAllCustomer(q Query, c *[]Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(CustomerDatabase).C(CustomerCollection).Find(q).All(c)
}

func (s *Store) FindOneCustomer(q Query, c *Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB(CustomerDatabase).C(CustomerCollection).Find(q).One(c)
}

func (s *Store) Close() {
	s.mongo.Close()

	s.mongo = nil
}
