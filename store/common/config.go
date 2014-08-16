package common

import "github.com/satisfeet/hoopoe/store/mongo"

// The Config type provides an unified interface to define storage settings for
// multiple storage engines.
type Config struct {
	Name   string
	Index  []string
	Unique []string
}

// Returns configuration converted for mongodb storage engine.
func (c Config) mongo() mongo.Config {
	return mongo.Config{
		Name:   c.Name,
		Index:  c.Index,
		Unique: c.Unique,
	}
}
