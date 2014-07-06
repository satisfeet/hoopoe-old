package store

import (
	"reflect"
	"strings"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Manager struct {
	collection *mgo.Collection
}

func (m *Manager) index(model interface{}) {
	v := reflect.ValueOf(model).Elem()

	var index, unique []string

	for i := 0; i < v.NumField(); i++ {
		t := v.Type().Field(i)

		if s := t.Tag.Get("store"); len(s) != 0 {
			if strings.Contains(s, "index") {
				index = append(index, t.Name)
			}
			if strings.Contains(s, "unique") {
				unique = append(unique, t.Name)
			}
		}
	}

	m.collection.EnsureIndex(mgo.Index{Key: index})
	m.collection.EnsureIndex(mgo.Index{Key: unique, Unique: true})
}

func (m *Manager) Find(models interface{}) error {
	return m.collection.Find(nil).All(models)
}

func (m *Manager) Create(model interface{}) error {
	m.index(model)

	return m.collection.Insert(model)
}

func (m *Manager) FindById(id string, model interface{}) error {
	m.index(model)

	return m.collection.FindId(bson.ObjectIdHex(id)).One(model)
}

func (m *Manager) UpdateById(id string, model interface{}) error {
	m.index(model)

	return m.collection.UpdateId(bson.ObjectIdHex(id), model)
}

func (m *Manager) DestroyById(id string) error {
	return m.collection.RemoveId(bson.ObjectIdHex(id))
}
