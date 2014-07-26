package store

import "github.com/satisfeet/hoopoe/store/mongodb"

func Open(s string) error {
	if err := mongodb.DefaultStore.Open(s); err != nil {
		return err
	}
	return nil
}

func Close() error {
	if err := mongodb.DefaultStore.Close(); err != nil {
		return err
	}
	return nil
}
