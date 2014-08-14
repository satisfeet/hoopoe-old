package store

import (
	"strings"

	"github.com/satisfeet/go-validation"
	"github.com/satisfeet/hoopoe/store/mongo"
	"github.com/satisfeet/hoopoe/utils"
)

type store struct {
	mongo *mongo.Store
}

func (s *store) Find(model interface{}) error {
	return s.mongo.Find(getName(model), nil, model)
}

func (s *store) FindOne(model interface{}) error {
	return s.mongo.FindId(getName(model), getId(model), model)
}

func (s *store) Insert(model interface{}) error {
	if err := validation.Validate(model); err != nil {
		return err
	}

	return s.mongo.Insert(getName(model), model)
}

func (s *store) Index(model interface{}) error {
	n := getName(model)
	i, u := getIndex(model)

	if len(i) > 0 {
		if err := s.mongo.Index(n, i); err != nil {
			return err
		}
	}
	if len(u) > 0 {
		return s.mongo.Unique(n, u)
	}

	return nil
}

func (s *store) Update(model interface{}) error {
	q := mongo.Query{}
	q.Id(getId(model))

	if err := validation.Validate(model); err != nil {
		return err
	}

	return s.mongo.Update(getName(model), q, model)
}

func (s *store) Remove(model interface{}) error {
	q := mongo.Query{}
	q.Id(getId(model))

	return s.mongo.Remove(getName(model), q)
}

func getId(model interface{}) interface{} {
	return utils.GetFieldValue(model, "Id")
}

func getName(model interface{}) string {
	return strings.ToLower(utils.GetTypeName(model)) + "s"
}

func getIndex(model interface{}) ([]string, []string) {
	u := []string{}
	i := []string{}

	for _, f := range utils.GetStructInfo(model) {
		if f.Unique {
			u = append(u, f.Name)
		}
		if f.Index {
			i = append(i, f.Name)
		}
	}

	return i, u
}
