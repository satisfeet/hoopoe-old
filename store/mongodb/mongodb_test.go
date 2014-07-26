package mongodb

import (
	"testing"

	"github.com/satisfeet/hoopoe/store/common"
	"gopkg.in/check.v1"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestMongodb(t *testing.T) {
	check.Suite(&QuerySuite{})
	check.Suite(&StoreSuite{
		Url: "localhost/test",
	})
	check.TestingT(t)
}

type QuerySuite struct{}

func (s *QuerySuite) TestId(c *check.C) {
	id := bson.NewObjectId()
	q1 := Query{}
	//q2 := Query{}

	c.Check(Query{}.Id(0), check.Equals, common.ErrBadQueryValue)
	c.Check(Query{}.Id(nil), check.Equals, common.ErrBadQueryValue)

	c.Check(Query{}.Id("abcd"), check.Equals, common.ErrBadQueryId)
	c.Check(Query{}.Id(bson.ObjectId("abcd")), check.Equals, common.ErrBadQueryId)

	c.Check(q1.Id(id.Hex()), check.IsNil)
	c.Check(q1.Id(id), check.IsNil)
	c.Check(q1["_id"], check.Equals, id)

	// TODO: needs fix
	//c.Check(q2["_id"], check.Equals, id)
}

type Sample struct {
	Id   bson.ObjectId `bson:"_id"`
	Name string
}

func (s Sample) Validate() error {
	return nil
}

type StoreSuite struct {
	Url     string
	Model   Sample
	Store   *Store
	Session *mgo.Session
}

func (s *StoreSuite) SetUpSuite(c *check.C) {
	Session, err := mgo.Dial(s.Url)
	c.Assert(err, check.IsNil)

	s.Session = Session
	s.Store = &Store{Session}
}

func (s *StoreSuite) SetUpTest(c *check.C) {
	s.Model = Sample{
		Id:   bson.NewObjectId(),
		Name: "foo",
	}

	c.Assert(s.Session.DB("").C("sample").Insert(&s.Model), check.IsNil)
}

func (s *StoreSuite) TestOpenAndClose(c *check.C) {
	store := &Store{}

	c.Assert(store.Open(s.Url), check.IsNil)
	c.Assert(store.Close(), check.IsNil)
}

func (s *StoreSuite) TestFind(c *check.C) {
	m := []Sample{}

	c.Assert(s.Store.Find(Query{}, &m), check.IsNil)
	c.Assert(m[0], check.DeepEquals, s.Model)
}

func (s *StoreSuite) TestFindOne(c *check.C) {
	m := Sample{}
	q := Query{}
	q.Id(s.Model.Id)

	c.Assert(s.Store.FindOne(q, &m), check.IsNil)
	c.Check(m, check.DeepEquals, s.Model)
}

func (s *StoreSuite) TestInsert(c *check.C) {
	m1 := Sample{
		Id:   bson.NewObjectId(),
		Name: "insert",
	}
	m2 := Sample{}
	q := Query{"_id": m1.Id}

	c.Assert(s.Store.Insert(&m1), check.IsNil)
	c.Check(s.Session.DB("").C("sample").Find(q).One(&m2), check.IsNil)
	c.Check(m1, check.DeepEquals, m2)
}

func (s *StoreSuite) TestUpdate(c *check.C) {
	s.Model.Name += "1234"
	m := Sample{}
	q := Query{"_id": s.Model.Id}

	c.Assert(s.Store.Update(s.Model), check.IsNil)
	c.Check(s.Session.DB("").C("sample").Find(q).One(&m), check.IsNil)
	c.Check(m, check.DeepEquals, m)
}

func (s *StoreSuite) TestRemove(c *check.C) {
	m := Sample{}
	q := Query{"_id": s.Model.Id}

	c.Assert(s.Store.Remove(s.Model), check.IsNil)
	c.Check(s.Session.DB("").C("sample").Find(q).One(&m), check.ErrorMatches, "not found")
}

func (s *StoreSuite) TearDownTest(c *check.C) {
	c.Assert(s.Session.DB("").C("sample").DropCollection(), check.IsNil)
}

func (s *StoreSuite) TearDownSuite(c *check.C) {
	s.Session.Close()
}
