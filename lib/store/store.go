package store

import (
    "labix.org/v2/mgo"

    "github.com/satisfeet/hoopoe/lib/conf"
)

var store struct {
    session *mgo.Session
    db      *mgo.Database
}

func Init() error {
    s, err := mgo.Dial(conf.Get("store")["host"])

    if err != nil {
        return err
    }

    store.session = s

    store.db = s.DB(conf.Get("store")["name"])

    return nil
}
