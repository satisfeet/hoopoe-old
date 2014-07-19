package store

import (
	"reflect"
	"testing"

	"gopkg.in/mgo.v2/bson"
)

var (
	queryIds = []string{
		bson.NewObjectId().Hex(),
		bson.NewObjectId().Hex(),
		bson.NewObjectId().Hex(),
		"1234567890123456789",
		"abcd",
		"",
	}
	querySearch = Query{
		"$or": []bson.M{
			bson.M{"foo": bson.RegEx{"bar", "i"}},
			bson.M{"baz": bson.RegEx{"bar", "i"}},
		},
	}
)

func TestQueryId(t *testing.T) {
	for _, v := range queryIds {
		q := Query{}
		q.Id(v)

		if bson.IsObjectIdHex(v) {
			if q["_id"] == nil {
				t.Error("Expected id to get set but was nil.\n")
			}
		} else {
			if q["_id"] != nil {
				t.Error("Expected id to not get set.\n")
			}
		}
	}
}

func TestQuerySearch(t *testing.T) {
	q1 := Query{}
	q2 := Query{}

	q1.Search("", []string{"foo", "baz"})
	q2.Search("bar", []string{"foo", "baz"})

	if q1["$or"] != nil {
		t.Errorf("Expect query to be empty but it was %v.\n", q1)
	}
	if !reflect.DeepEqual(q2, querySearch) {
		t.Errorf("Expected query to equal search query but it was %v.\n", q2)
	}
}
