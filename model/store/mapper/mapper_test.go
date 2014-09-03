package mapper

import (
	"database/sql"
	"reflect"
	"testing"

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
	m := NewMapper(&s.people, s.columns)

	for i := 0; i < 3; i++ {
		s.scan(m.Params()...)

		err := m.Scan()
		c.Assert(err, check.IsNil)
	}

	c.Check(s.people, check.DeepEquals, []person{s.result, s.result, s.result})
}

func (s *MapperSuite) TestStruct(c *check.C) {
	m := NewMapper(&s.person, s.columns)

	s.scan(m.Params()...)

	err := m.Scan()
	c.Assert(err, check.IsNil)

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
	reflect.ValueOf(params[0]).Elem().Set(reflect.ValueOf(sql.RawBytes("Joe")))
	reflect.ValueOf(params[1]).Elem().Set(reflect.ValueOf(sql.RawBytes("Good,Nice")))
	reflect.ValueOf(params[2]).Elem().Set(reflect.ValueOf(sql.RawBytes("Some City")))
	reflect.ValueOf(params[3]).Elem().Set(reflect.ValueOf(sql.RawBytes("12345")))
}
