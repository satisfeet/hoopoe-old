package store

import (
	"testing"

	"gopkg.in/check.v1"
)

func TestSuite(t *testing.T) {
	check.Suite(&Suite{})
	check.TestingT(t)
}

type Suite struct{}

type model struct {
	Name string `store:"index"`
}
