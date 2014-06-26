package store

import (
    "labix.org/v2/mgo"

    "github.com/satisfeet/hoopoe/lib/conf"
)

type Store struct {
    session *mgo.Session
    db      *mgo.Database
}

func New() *Store {
    return &Store{}
}

func (s *Store) Open(c *conf.Conf) error {
    var err error

    s.session, err = mgo.Dial(c.Store["host"])

    if err != nil {
        return err
    }

    s.db = s.session.DB(c.Store["name"])

    return nil
}
