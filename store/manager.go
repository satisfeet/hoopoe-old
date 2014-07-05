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

type Model interface {
	Id() bson.ObjectId
	NewId()
}

func (m *Manager) index(model Model) {
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

func (m *Manager) Create(model Model) error {
	m.index(model)

	if !model.Id().Valid() {
		model.NewId()
	}

	return m.collection.Insert(model)
}

func (m *Manager) Update(model Model) error {
	m.index(model)

	return m.collection.UpdateId(model.Id(), model)
}

func (m *Manager) Destroy(model Model) error {
	m.index(model)

	return m.collection.RemoveId(model.Id())
}

func (m *Manager) Find(query Query, models interface{}) error {
	return m.collection.Find(query).All(models)
}

func (m *Manager) FindOne(query Query, model Model) error {
	m.index(model)

	return m.collection.Find(query).One(model)
}
