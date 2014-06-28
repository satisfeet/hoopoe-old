package store

import (
    "labix.org/v2/mgo"
)

var (
    session *mgo.Session
    db      *mgo.Database
)

func Open(config map[string]string) error {
    var err error

    session, err = mgo.Dial(config["mongo"])

    if err != nil {
        return err
    }

    db = session.DB("")

    return nil
}
