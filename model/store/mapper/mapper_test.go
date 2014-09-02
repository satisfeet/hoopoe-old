package mapper

import (
	"database/sql"
	"testing"

	"github.com/satisfeet/hoopoe/model/utils"
	"gopkg.in/check.v1"
)

func TestMapper(t *testing.T) {
	check.Suite(&MapperSuite{
		columns: []string{
			"name",
			"tags",
			"city",
			"code",
		},
	})
	check.TestingT(t)
}

type MapperSuite struct {
	columns []string
	person  person
	people  []person
	result  person
}

func (s *MapperSuite) SetUpTest(c *check.C) {
	s.person = person{}
	s.people = []person{}

	s.result = person{
		Name: "Joe",
		Tags: []string{
			"Good",
			"Nice",
		},
		Address: address{
			City: "Some City",
			Code: 12345,
		},
	}
}

func (s *MapperSuite) TestSlice(c *check.C) {
	m := NewMapper(&s.people)
	m.SetColumns(s.columns)

	for i := 0; i < 3; i++ {
		src := m.NewSource()
		s.scan(src.Params()...)

		err := m.MapSource(src)
		c.Assert(err, check.IsNil)
	}

	c.Check(s.people, check.DeepEquals, []person{s.result, s.result, s.result})
}

func (s *MapperSuite) TestStruct(c *check.C) {
	m := NewMapper(&s.person)
	m.SetColumns(s.columns)

	src := m.NewSource()
	s.scan(src.Params()...)
	m.MapSource(src)

	c.Check(s.person, check.DeepEquals, s.result)
}

type person struct {
	Name    string
	Tags    []string
	Address address
}

type address struct {
	Street string
	City   string
	Code   int
}

func (s *MapperSuite) scan(params ...interface{}) {
	utils.SetValue(params[0], sql.RawBytes("Joe"))
	utils.SetValue(params[1], sql.RawBytes("Good,Nice"))
	utils.SetValue(params[2], sql.RawBytes("Some City"))
	utils.SetValue(params[3], sql.RawBytes("12345"))
}
