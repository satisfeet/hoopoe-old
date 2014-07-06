package store

import (
	"reflect"
	"strings"

	"labix.org/v2/mgo"
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

func (m *Manager) Create(model interface{}) error {
	m.index(model)

	return m.collection.Insert(model)
}

func (m *Manager) Update(query Query, model interface{}) error {
	m.index(model)

	return m.collection.Update(query, model)
}

func (m *Manager) Destroy(query Query) error {
	return m.collection.Remove(query)
}

func (m *Manager) Find(query Query, models interface{}) error {
	return m.collection.Find(query).All(models)
}

func (m *Manager) FindOne(query Query, model interface{}) error {
	m.index(model)

	return m.collection.Find(query).One(model)
}
