package mongo

// The Config type defines settings for the mongodb engine.
type Config struct {
	Name   string
	Index  []string
	Unique []string
}
