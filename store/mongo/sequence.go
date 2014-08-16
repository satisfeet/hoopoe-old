package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Mgo operation used to update a value in sequences.
var incOp = bson.M{"$inc": bson.M{"value": 1}}

// Mgo operation used to reset a value in sequences.
var resetOp = bson.M{"value": 0}

// The Sequence type can be used as generator for auto incremented id values as
// these are not supported by mongodb out of the box.
type Sequence struct {
	name       string
	collection *mgo.Collection
}

// Returns an initialized Sequence.
func NewSequence(c Config, s *Session) *Sequence {
	return &Sequence{
		name:       c.Name,
		collection: s.database.C("sequences"),
	}
}

// Returns increased sequence value and error of operation.
func (s *Sequence) New() (int, error) {
	m := make(map[string]int)

	_, err := s.collection.FindId(s.name).Apply(mgo.Change{
		Update:    incOp,
		ReturnNew: true,
	}, &m)

	return m["value"], err
}

// Returns error of reset operation.
func (s *Sequence) Reset() error {
	_, err := s.collection.UpsertId(s.name, resetOp)

	return err
}
