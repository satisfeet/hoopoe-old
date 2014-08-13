package mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/utils"
)

var ErrBadId = errors.New("bad id value")

type Store struct {
	session  *mgo.Session
	database *mgo.Database
}

func NewStore(s *mgo.Session) *Store {
	return &Store{
		session:  s,
		database: s.DB(""),
	}
}

func (s *Store) Dial(url string) error {
	sess, err := mgo.Dial(url)

	if err != nil {
		return err
	}

	s.session = sess
	s.database = sess.DB("")

	return nil
}

func (s *Store) Close() error {
	s.session.Close()

	s.session = nil
	s.database = nil

	return nil
}

func (s *Store) Find(name string, q Query, value interface{}) error {
	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).Find(q).All(value)
}

func (s *Store) FindId(name string, id interface{}, value interface{}) error {
	oid, err := castId(id)

	if err != nil {
		return err
	}

	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).FindId(oid).One(value)
}

func (s *Store) FindOne(name string, q Query, value interface{}) error {
	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).Find(q).One(value)
}

func (s *Store) Insert(name string, value interface{}) error {
	sess := s.clone()
	defer sess.Close()

	if id := utils.GetFieldValue(value, "Id").(bson.ObjectId); !id.Valid() {
		utils.SetFieldValue(value, "Id", bson.NewObjectId())
	}

	return s.collection(name).With(sess).Insert(value)
}

func (s *Store) Update(name string, q Query, value interface{}) error {
	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).Update(q, value)
}

func (s *Store) UpdateId(name string, id interface{}, value interface{}) error {
	oid, err := castId(id)

	if err != nil {
		return err
	}

	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).UpdateId(oid, value)
}

func (s *Store) Remove(name string, q Query) error {
	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).Remove(q)
}

func (s *Store) RemoveId(name string, id interface{}) error {
	oid, err := castId(id)

	if err != nil {
		return err
	}

	sess := s.clone()
	defer sess.Close()

	return s.collection(name).With(sess).RemoveId(oid)
}

func (s *Store) RemoveAll(name string, q Query) error {
	sess := s.clone()
	defer sess.Close()

	_, err := s.collection(name).With(sess).RemoveAll(q)

	return err
}

func (s *Store) CreateFile(name string) (*File, error) {
	f, err := s.filesystem(name).Create("")

	if err != nil {
		return nil, err
	}

	id, ok := f.Id().(bson.ObjectId)
	if !ok || !id.Valid() {
		id = bson.NewObjectId()

		f.SetId(id)
	}

	return &File{id, f}, nil
}

func (s *Store) OpenFileId(name string, id interface{}) (*File, error) {
	oid, err := castId(id)

	if err != nil {
		return nil, err
	}

	f, err := s.filesystem(name).OpenId(oid)

	if err != nil {
		return nil, err
	}

	return &File{oid, f}, nil
}

func (s *Store) RemoveFileId(name string, id interface{}) error {
	oid, err := castId(id)

	if err != nil {
		return err
	}

	return s.filesystem(name).RemoveId(oid)
}

func (s *Store) clone() *mgo.Session {
	return s.session.Clone()
}

func (s *Store) filesystem(name string) *mgo.GridFS {
	return s.database.GridFS(name)
}

func (s *Store) collection(name string) *mgo.Collection {
	return s.database.C(name)
}

func castId(id interface{}) (bson.ObjectId, error) {
	var oid bson.ObjectId

	switch t := id.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			oid = bson.ObjectIdHex(t)
		}
	case bson.ObjectId:
		oid = t
	}

	if oid.Valid() {
		return oid, nil
	}

	return oid, ErrBadId
}

func IdFromString(id string) bson.ObjectId {
	var oid bson.ObjectId

	if bson.IsObjectIdHex(id) {
		oid = bson.ObjectIdHex(id)
	}

	return oid
}
