package store

import (
	"testing"

	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2/bson"
)

func TestSuite(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct {
	id bson.ObjectId
}

type model struct {
	Name string `store:"index"`
}

func (s *Suite) SetUpTest(c *check.C) {
	s.id = bson.NewObjectId()
}
