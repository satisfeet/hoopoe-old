package store

import (
    "labix.org/v2/mgo"
)

var (
    session *mgo.Session
    db      *mgo.Database
)

func Open(c map[string]string) error {
    var err error

    session, err = mgo.Dial(c["host"])

    if err != nil {
        return err
    }

    db = session.DB(c["name"])

    return nil
}
