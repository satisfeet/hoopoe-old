package store

import (
	"strings"
	"testing"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	name = "testers"

	modelsWithId = []model{
		model{
			Id:   bson.NewObjectId(),
			Text: "I have an id",
		},
		model{
			Id:   bson.NewObjectId(),
			Text: "I also have an id",
		},
	}
	modelsWithoutId = []model{
		model{
			Text: "I have no id",
		},
		model{
			Text: "I also have no id",
		},
	}
)

type (
	model struct {
		Id   bson.ObjectId `bson:"_id"`
		Text string
	}
)

func TestStoreInsert(t *testing.T) {
	for _, v := range modelsWithoutId {
		setup(nil)

		if err := NewStore(name).Insert(&v); err != nil {
			t.Error(err)
		}
		if i := count(); i != 1 {
			t.Errorf("Expected to find one document but found %d.\n", i)
		}

		teardown()
	}
}

func TestStoreUpdate(t *testing.T) {
	for _, v := range modelsWithId {
		setup(&v)
		v.Text += "1234?"

		if err := NewStore(name).Update(Query{"_id": v.Id}, &v); err != nil {
			t.Error(err)
		}
		mongo.DB(Database).C(name).FindId(v.Id).One(&v)
		if !strings.HasSuffix(v.Text, "1234?") {
			t.Error("Expected document to be updated.\n")
		}

		teardown()
	}
}

func TestStoreRemove(t *testing.T) {
	for _, v := range modelsWithId {
		setup(v)

		if err := NewStore(name).Remove(Query{"_id": v.Id}); err != nil {
			t.Error(err)
		}
		if i := count(); i != 0 {
			t.Errorf("Expected to have no documents but found %d.\n", i)
		}

		teardown()
	}
}

func count() int {
	i, _ := mongo.DB("").C(name).Find(nil).Count()

	return i
}

func setup(v interface{}) {
	mongo, _ = mgo.Dial("localhost/test")
	mongo.DB("").C(name).DropCollection()

	if v != nil {
		mongo.DB("").C(name).Insert(v)
	}
}

func teardown() {
	mongo.DB("").C(name).DropCollection()
	mongo.Close()
}
