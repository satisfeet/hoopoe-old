package mongo

import "gopkg.in/mgo.v2"

// The Store type exposes common operation for interaction with mongodb.
type Store struct {
	config   Config
	database *mgo.Database
	Sequence *Sequence
}

// Returns an initialized mongo Store.
func NewStore(c Config, s *Session) *Store {
	return &Store{
		config:   c,
		database: s.database,
		Sequence: NewSequence(c, s),
	}
}

// Creates the defined indices if necessary.
func (s *Store) Index() error {
	if k := s.config.Index; len(k) > 0 {
		return s.collection().EnsureIndexKey(k...)
	}
	if k := s.config.Unique; len(k) > 0 {
		i := mgo.Index{Key: k, Unique: true}

		return s.collection().EnsureIndex(i)
	}

	return nil
}

// Maps documents matching query conditions onto models.
func (s *Store) FindAll(q *Query, models interface{}) error {
	if err := q.Err(); err != nil {
		return err
	}

	return s.collection().Find(q.query).All(models)
}

// Maps document matching query condition onto model.
func (s *Store) FindOne(q *Query, model interface{}) error {
	if err := q.Err(); err != nil {
		return err
	}

	return s.collection().Find(q.query).One(model)
}

// Inserts new model into storage.
func (s *Store) Insert(model interface{}) error {
	return s.collection().Insert(model)
}

// Updates model from storage.
func (s *Store) Update(q *Query, model interface{}) error {
	return s.collection().Update(q.query, model)
}

// Removes model from storage.
func (s *Store) Remove(q *Query) error {
	return s.collection().Remove(q.query)
}

// Returns mgo collection defined per config.
func (s *Store) collection() *mgo.Collection {
	return s.database.C(s.config.Name)
}
