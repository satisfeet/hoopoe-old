package store

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Manager struct {
	collection *mgo.Collection
}

type Model interface {
	Id() bson.ObjectId
}

func (m *Manager) Create(model Model) error {
	return m.collection.Insert(model)
}

func (m *Manager) Update(model Model) error {
	return m.collection.UpdateId(model.Id(), model)
}

func (m *Manager) Destroy(model Model) error {
	return m.collection.RemoveId(model.Id())
}

func (m *Manager) FindOne(query Query, model Model) error {
	return m.collection.Find(query).One(model)
}

func (m *Manager) Find(query Query, models interface{}) error {
	return m.collection.Find(query).All(models)
}
