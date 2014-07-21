package store

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var (
	ErrNotConnected = errors.New("not connected")
)

const Database = ""

type Session struct {
	mongo *mgo.Session
}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) Open(u string) error {
	var err error
	s.mongo, err = mgo.Dial(u)
	return err
}

func (s *Session) Mongo() *mgo.Session {
	if s.mongo == nil {
		panic(ErrNotConnected)
	}

	return s.mongo.Clone()
}

func (s *Session) Close() {
	if s.mongo == nil {
		panic(ErrNotConnected)
	}

	s.mongo.Close()
	s.mongo = nil
}
