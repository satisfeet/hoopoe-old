package store

import (
	"errors"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

var ErrBadId = errors.New("bad id")

type Model interface {
	Validate() error
}

type store struct {
	session  *mgo.Session
	database *mgo.Database
}

func (s *store) files(model interface{}) *mgo.GridFS {
	return s.database.GridFS(getName(model))
}

func (s *store) collection(model interface{}) *mgo.Collection {
	return s.database.C(getName(model))
}

func (s *store) Find(models interface{}) error {
	c := s.session.Clone()
	defer c.Close()

	return s.collection(models).With(c).Find(nil).All(models)
}

func (s *store) FindOne(model Model) error {
	c := s.session.Clone()
	defer c.Close()

	id := getId(model)

	if !id.Valid() {
		return ErrBadId
	}

	return s.collection(model).With(c).FindId(id).One(model)
}

func (s *store) Insert(model Model) error {
	c := s.session.Clone()
	defer c.Close()

	if id := getId(model); !id.Valid() {
		setId(model, bson.NewObjectId())
	}

	if err := model.Validate(); err != nil {
		return err
	}

	return s.collection(model).With(c).Insert(model)
}

func (s *store) Update(model Model) error {
	c := s.session.Clone()
	defer c.Close()

	id := getId(model)

	if !id.Valid() {
		return ErrBadId
	}

	if err := model.Validate(); err != nil {
		return err
	}

	return s.collection(model).With(c).UpdateId(id, model)
}

func (s *store) Remove(model Model) error {
	c := s.session.Clone()
	defer c.Close()

	id := getId(model)

	if !id.Valid() {
		return ErrBadId
	}

	return s.collection(model).With(c).RemoveId(id)
}

func setId(model interface{}, id bson.ObjectId) {
	utils.SetFieldValue(model, "Id", id)
}

func getId(model interface{}) bson.ObjectId {
	return utils.GetFieldValue(model, "Id").(bson.ObjectId)
}

func getName(model interface{}) string {
	return strings.ToLower(utils.GetTypeName(model)) + "s"
}

func IdFromString(id string) bson.ObjectId {
	var oid bson.ObjectId

	if bson.IsObjectIdHex(id) {
		oid = bson.ObjectIdHex(id)
	}

	return oid
}
