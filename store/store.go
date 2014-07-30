package store

import "github.com/satisfeet/hoopoe/store/mongo"

var DefaultMongo = &mongo.Store{}

func Dial(s string) error {
	if err := DefaultMongo.Dial(s); err != nil {
		return err
	}
	return nil
}

func Close() error {
	if err := DefaultMongo.Close(); err != nil {
		return err
	}
	return nil
}
