package store

import (
	"testing"

	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"
)

type Suite struct {
	Customers []Customer
}

func (s *Suite) SetUp() {
	err := Open(map[string]string{
		"mongo": "localhost/test",
	})

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

	err = db.C("customers").Insert(
		s.Customers[0],
		s.Customers[1],
	)

	So(err, ShouldBeNil)
}

func TestCustomersCreate(t *testing.T) {
	s := &Suite{}

	Convey("CustomersCreate", t, func() {
		s.SetUp()

		c := Customer{
			Name:  "Edison Trent",
			Email: "edison@liberty.si",
			Address: CustomerAddress{
				City: "Leeds",
			},
		}

		err := CustomersCreate(&c)

		Convey("Should have no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("Should have valid id", func() {
			So(c.Id.Valid(), ShouldEqual, true)
		})
		Convey("Should have created customer", func() {
			r := Customer{}

			err := db.C("customers").FindId(c.Id).One(&r)

			So(err, ShouldBeNil)
			So(&r, ShouldResemble, &c)
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersUpdate(t *testing.T) {
	s := &Suite{}

	Convey("CustomersUpdate", t, func() {
		s.SetUp()

		Convey("With Customer", func() {
			s.Customers[0].Address.City = "New City"

			err := CustomersUpdate(&s.Customers[0])

			Convey("Should have no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should have updated customer", func() {
				r := &Customer{}

				err := db.C("customers").FindId(s.Customers[0].Id).One(&r)

				So(err, ShouldBeNil)
				So(r, ShouldResemble, &s.Customers[0])
				So(r.Address.City, ShouldEqual, "New City")
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersRemove(t *testing.T) {
	s := &Suite{}

	Convey("CustomersRemove", t, func() {
		s.SetUp()

		Convey("With Customer", func() {
			err := CustomersRemove(&s.Customers[0])

			Convey("Should have no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should have removed customer", func() {
				i, err := db.C("customers").FindId(s.Customers[0].Id).Count()

				So(err, ShouldBeNil)
				So(i, ShouldEqual, 0)
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersFindOne(t *testing.T) {
	s := &Suite{}

	Convey("CustomersFindOne", t, func() {
		s.SetUp()

		Convey("With id Query", func() {
			c, err := CustomersFindOne(Query{
				"id": s.Customers[1].Id.Hex(),
			})

			Convey("Should have no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should have result to be customer", func() {
				So(c, ShouldResemble, s.Customers[1])
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func TestCustomersFindAll(t *testing.T) {
	s := &Suite{}

	Convey("CustomersFindAll", t, func() {
		s.SetUp()

		Convey("With empty Query", func() {
			c, err := CustomersFindAll(Query{})

			Convey("Should have no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should have result to be customers", func() {
				So(c, ShouldHaveSameTypeAs, []Customer{})

				Convey("With length two", func() {
					So(len(c), ShouldEqual, 2)
				})
				Convey("With all customers", func() {
					So(c[0], ShouldResemble, s.Customers[0])
					So(c[1], ShouldResemble, s.Customers[1])
				})
			})
		})
		Convey("With search Query", func() {
			c, err := CustomersFindAll(Query{
				"search": "Haci",
			})

			Convey("Should have no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should have result to be customers", func() {
				So(c, ShouldHaveSameTypeAs, []Customer{})

				Convey("With length one", func() {
					So(len(c), ShouldEqual, 1)
				})
				Convey("With one customer", func() {
					So(c[0], ShouldResemble, s.Customers[0])
				})
			})
		})

		Reset(func() {
			s.TearDown()
		})
	})
}

func (s *Suite) TearDown() {
	_, err := db.C("customers").RemoveAll(bson.M{})

	So(err, ShouldBeNil)

	Close()
}
