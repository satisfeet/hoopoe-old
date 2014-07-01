package customers

import (
	"testing"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"
)

type Suite struct {
	Customers  []Customer
	Session    *mgo.Session
	Collection *mgo.Collection
}

func (s *Suite) SetUp() {
	var err error

	s.Session, err = mgo.Dial("localhost/test")

	So(err, ShouldBeNil)

	s.Customers = []Customer{
		Customer{
			Id:    bson.NewObjectId(),
			Name:  "Haci Erdal",
			Email: "haci95@hotmail.com",
			Address: CustomerAddress{
				City: "Berlin",
			},
		},
		Customer{
			Id:      bson.NewObjectId(),
			Name:    "Bodo Kaiser",
			Email:   "i@bodokaiser.io",
			Company: "satisfeet",
			Address: CustomerAddress{
				Street: "Geiserichstraße 3",
				City:   "Berlin",
				Zip:    12105,
			},
		},
	}

	s.Collection = s.Session.DB("").C("customers")

	Setup(s.Session.DB(""))
}

func (s *Suite) Insert() {
	err := s.Collection.Insert(s.Customers[0], s.Customers[1])

	So(err, ShouldBeNil)
}

func (s *Suite) Remove() {
	_, err := s.Collection.RemoveAll([]bson.M{})

	So(err, ShouldBeNil)
}

func (s *Suite) TearDown() {
	s.Session.Close()
}

func TestCreate(test *testing.T) {
	s := &Suite{}

	Convey(`"Create"`, test, func() {
		s.SetUp()
		s.Remove()

		Convey(`With "Customer"`, func() {
			c := Customer{
				Name:  "Edison Trent",
				Email: "edison@liberty.si",
				Address: CustomerAddress{
					City: "Leeds",
				},
			}

			err := Create(&c)

			Convey(`Should have "err" equal "nil"`, func() {
				So(err, ShouldBeNil)
			})
			Convey(`Should have set "Customer.Id"`, func() {
				So(c.Id, ShouldHaveSameTypeAs, bson.NewObjectId())
			})
			Convey(`Should be saved to "Collection"`, func() {
				r := Customer{}

				err := s.Collection.FindId(c.Id).One(&r)

				So(err, ShouldBeNil)
				So(&r, ShouldResemble, &c)
			})
		})

		Reset(func() {
			s.Remove()
			s.TearDown()
		})
	})
}

func TestUpdate(test *testing.T) {
	s := &Suite{}

	Convey(`"Update"`, test, func() {
		s.SetUp()
		s.Remove()
		s.Insert()

		Convey(`With "Customer"`, func() {
			s.Customers[0].Address.City = "New City"

			err := Update(&s.Customers[0])

			Convey(`Should have "err" equal "nil"`, func() {
				So(err, ShouldBeNil)
			})
			Convey(`Should be saved to "Collection"`, func() {
				r := &Customer{}

				err := s.Collection.FindId(s.Customers[0].Id).One(&r)

				So(err, ShouldBeNil)
				So(r, ShouldResemble, &s.Customers[0])
				So(r.Address.City, ShouldEqual, "New City")
			})
		})

		Reset(func() {
			s.Remove()
			s.TearDown()
		})
	})
}

func TestRemove(test *testing.T) {
	s := &Suite{}

	Convey(`"Remove"`, test, func() {
		s.SetUp()
		s.Remove()
		s.Insert()

		Convey(`With "Customer"`, func() {
			err := Remove(&s.Customers[0])

			Convey(`Should have "err" equal "nil"`, func() {
				So(err, ShouldBeNil)
			})
			Convey(`Should be removed from "Collection"`, func() {
				r, err := s.Collection.FindId(s.Customers[0].Id).Count()

				So(err, ShouldBeNil)
				So(r, ShouldEqual, 0)
			})
		})

		Reset(func() {
			s.Remove()
			s.TearDown()
		})
	})
}

func TestFindAll(test *testing.T) {
	s := &Suite{}

	Convey(`"FindAll"`, test, func() {
		s.SetUp()
		s.Remove()
		s.Insert()

		Convey(`With empty "Query"`, func() {
			c, err := FindAll(Query{})

			Convey(`Should have "err" equal "nil"`, func() {
				So(err, ShouldBeNil)
			})
			Convey(`Should have "result" include all models`, func() {
				So(len(c), ShouldEqual, 2)

				So(c[0], ShouldResemble, s.Customers[0])
				So(c[1], ShouldResemble, s.Customers[1])
			})
		})
		Convey(`With a search "Query"`, func() {
			c, err := FindAll(Query{"search": "Haci"})

			Convey(`Should have "err" equal "nil"`, func() {
				So(err, ShouldBeNil)
			})
			Convey(`Should have "result" include first model`, func() {
				So(len(c), ShouldEqual, 1)

				So(c[0], ShouldResemble, s.Customers[0])
			})
		})

		Reset(func() {
			s.Remove()
			s.TearDown()
		})
	})
}

func TestFindOne(test *testing.T) {
	s := &Suite{}

	Convey(`"FindOne"`, test, func() {
		s.SetUp()
		s.Remove()
		s.Insert()

		Convey(`With id "Query"`, func() {
			c, err := FindOne(Query{"id": s.Customers[1].Id.Hex()})

			Convey(`Should have "err" equal "nil"`, func() {
				So(err, ShouldBeNil)
			})
			Convey(`Should have "result" include second model`, func() {
				So(c, ShouldResemble, s.Customers[1])
			})
		})

		Reset(func() {
			s.Remove()
			s.TearDown()
		})
	})
}
