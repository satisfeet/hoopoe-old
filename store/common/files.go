package common

import (
	"io"

	"github.com/satisfeet/hoopoe/store/mongo"
)

type Files struct {
	mongo *mongo.Store
}

func NewFiles(c Config, s *Session) *Files {
	return &Files{
		mongo: mongo.NewStore(c.mongo(), s.mongo),
	}
}

func (fs *Files) New(id interface{}) (io.ReadWriteCloser, error) {
	return fs.mongo.FileSystem.New(id)
}

func (fs *Files) Open(id interface{}) (io.ReadWriteCloser, error) {
	return fs.mongo.FileSystem.Open(id)
}

func (fs *Files) Remove(id interface{}) error {
	return fs.mongo.FileSystem.Remove(id)
}
