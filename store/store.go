package store

import "github.com/satisfeet/hoopoe/store/mongo"

// Default mongo store to use.
var DefaultMongo = &mongo.Store{}

// Opens all global data stores.
func Open(s string) error {
	if err := DefaultMongo.Open(s); err != nil {
		return err
	}
	return nil
}

// Closes all global data stores.
func Close() error {
	if err := DefaultMongo.Close(); err != nil {
		return err
	}
	return nil
}
