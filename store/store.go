package store

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

var ErrBadId = errors.New("bad id")

type Model interface {
	Validate() error
}

type storeInfo struct {
	Name   string
	Index  []string
	Unique []string
}

type store struct {
	info     storeInfo
	session  *mgo.Session
	database *mgo.Database
}

func (s *store) files() *mgo.GridFS {
	return s.database.GridFS(s.info.Name)
}

func (s *store) collection() *mgo.Collection {
	return s.database.C(s.info.Name)
}

func (s *store) Index() error {
	if i := s.info.Index; len(i) > 0 {
		if err := s.collection().EnsureIndexKey(i...); err != nil {
			return err
		}
	}
	if i := s.info.Unique; len(i) > 0 {
		err := s.collection().EnsureIndex(mgo.Index{
			Key:    i,
			Unique: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *store) Find(models interface{}) error {
	c := s.session.Clone()
	defer c.Close()

	return s.collection().With(c).Find(nil).All(models)
}

func (s *store) FindOne(model Model) error {
	c := s.session.Clone()
	defer c.Close()

	id := getId(model)

	if !id.Valid() {
		return ErrBadId
	}

	return s.collection().With(c).FindId(id).One(model)
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

	return s.collection().With(c).Insert(model)
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

	return s.collection().With(c).UpdateId(id, model)
}

func (s *store) Remove(model Model) error {
	c := s.session.Clone()
	defer c.Close()

	id := getId(model)

	if !id.Valid() {
		return ErrBadId
	}

	return s.collection().With(c).RemoveId(id)
}

func setId(model interface{}, id bson.ObjectId) {
	utils.SetFieldValue(model, "Id", id)
}

func getId(model interface{}) bson.ObjectId {
	return utils.GetFieldValue(model, "Id").(bson.ObjectId)
}

func IdFromString(id string) bson.ObjectId {
	var oid bson.ObjectId

	if bson.IsObjectIdHex(id) {
		oid = bson.ObjectIdHex(id)
	}

	return oid
}
