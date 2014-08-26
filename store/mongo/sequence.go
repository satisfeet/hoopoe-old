package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var sequenceName = "sequences"

// Mgo operation used to update a value in sequences.
var incOp = bson.M{"$inc": bson.M{"value": 1}}

// Mgo operation used to reset a value in sequences.
var resetOp = bson.M{"value": 0}

// The Sequence type can be used as generator for auto incremented id values as
// these are not supported by mongodb out of the box.
type Sequence struct {
	config  Config
	session *Session
}

// Returns increased sequence value and error of operation.
func (s *Sequence) New() (int, error) {
	m := make(map[string]int)

	_, err := s.sequence().FindId(s.config.Name).Apply(mgo.Change{
		Update:    incOp,
		ReturnNew: true,
	}, &m)

	return m["value"], err
}

// Returns error of reset operation.
func (s *Sequence) Reset() error {
	_, err := s.sequence().UpsertId(s.config.Name, resetOp)

	return err
}

func (s *Sequence) sequence() *mgo.Collection {
	return s.session.database.C(sequenceName)
}
