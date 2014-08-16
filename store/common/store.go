package common

import "github.com/satisfeet/hoopoe/store/mongo"

// The Store type is provides a common interface for store interactions. Its
// intended to be used as embedded type.
type Store struct {
	mongo *mongo.Store
}

// Returns an initialized Store.
func NewStore(c Config, s *Session) *Store {
	return &Store{
		mongo: mongo.NewStore(c.mongo(), s.mongo),
	}
}

// Ensures Index on all storage engines.
func (s *Store) Index() error {
	return s.mongo.Index()
}

// Executes an operation which maps all documents matching conditions defined in
// query onto the provided models interface.
func (s *Store) FindAll(q query, models interface{}) error {
	return s.mongo.FindAll(q.mongo(), models)
}

// Executes an operation which maps one document matching conditions defined in
// query onto the provided model interface.
func (s *Store) FindOne(q query, model interface{}) error {
	return s.mongo.FindOne(q.mongo(), model)
}

// Inserts a validatable model into storage.
func (s *Store) Insert(model validatable) error {
	if err := model.Validate(); err != nil {
		return err
	}

	return s.mongo.Insert(model)
}

// Updates a validatable model from storage.
func (s *Store) Update(q query, model validatable) error {
	if err := model.Validate(); err != nil {
		return err
	}

	return s.mongo.Update(q.mongo(), model)
}

// Removes one document matching conditions from storage.
func (s *Store) Remove(q query) error {
	return s.mongo.Remove(q.mongo())
}

// Exposes the sequence manager from mongo package.
func (s *Store) Sequence() *mongo.Sequence {
	return s.mongo.Sequence
}

// The query interface is an helper to accept types with embedded Query in store
// actions.
type query interface {
	Id(interface{})

	mongo() *mongo.Query
}

// The validatable interface is an helper to check if a model provides a
// validation method.
type validatable interface {
	Validate() error
}
