package customers

import (
    "testing"

    "labix.org/v2/mgo"
    "labix.org/v2/mgo/bson"

    . "github.com/smartystreets/goconvey/convey"
)

type Suite struct {
    Models	[]Customer
    Session	*mgo.Session
}

func TestFindAll(test *testing.T) {
    s := &Suite{}

    Convey(`"FindAll"`, test, func() {
        s.SetUp()
        s.Remove()
        s.Insert()

        Convey(`With empty "Query"`, func() {
            q := &Query{}

            result, err := FindAll(q)

            Convey(`Should have "err" equal "nil"`, func() {
                So(err, ShouldBeNil)
            })
            Convey(`Should have "result" include all models`, func() {
                So(len(result), ShouldEqual, 2)

                So(result[0], ShouldResemble, s.Models[0])
                So(result[1], ShouldResemble, s.Models[1])
            })
        })
        Convey(`With a search "Query"`, func() {
            q := &Query{}
            q.Search("Haci")

            result, err := FindAll(q)

            Convey(`Should have "err" equal "nil"`, func() {
                So(err, ShouldBeNil)
            })
            Convey(`Should have "result" include first model`, func() {
                So(len(result), ShouldEqual, 1)

                So(result[0], ShouldResemble, s.Models[0])
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
            q := &Query{}
            q.Id(s.Models[1].Id.Hex())

            result, err := FindOne(q)

            Convey(`Should have "err" equal "nil"`, func() {
                So(err, ShouldBeNil)
            })
            Convey(`Should have "result" include second model`, func() {
                So(result, ShouldResemble, s.Models[1])
            })
        })

        Reset(func() {
            s.Remove()
            s.TearDown()
        })
    })
}

func (s *Suite) SetUp() {
    sess, err := mgo.Dial("localhost/test")

    So(err, ShouldBeNil)

    s.Models = append(s.Models, Customer{
        Id:			bson.NewObjectId(),
        Name:		"Haci Erdal",
        Email:		"haci95@hotmail.com",
        Address:	CustomerAddress{
            City:		"Berlin",
        },
    })
    s.Models = append(s.Models, Customer{
        Id:			bson.NewObjectId(),
        Name:		"Bodo Kaiser",
        Email:		"i@bodokaiser.io",
        Company:	"satisfeet",
        Address:	CustomerAddress{
            Street:		"Geiserichstraße 3",
            City:		"Berlin",
            Zip:		12105,
        },
    })
    s.Session = sess

    Setup(sess)
}

func (s *Suite) Insert() {
    err := s.Session.DB("").C("customers").Insert(
        s.Models[0],
        s.Models[1],
    )

    So(err, ShouldBeNil)
}

func (s *Suite) Remove() {
    _, err := s.Session.DB("").C("customers").RemoveAll([]bson.M{})

    So(err, ShouldBeNil)
}

func (s *Suite) TearDown() {
    s.Session.Close()
}
