package mongo

import (
	"io"

	"gopkg.in/mgo.v2"
)

type FileSystem struct {
	config  Config
	session *Session
}

func (fs *FileSystem) New(id interface{}) (io.ReadWriteCloser, error) {
	f, err := fs.files().Create("")

	if id != nil {
		f.SetId(id)
	}

	return f, err
}

func (fs *FileSystem) Open(id interface{}) (io.ReadWriteCloser, error) {
	return fs.files().OpenId(id)
}

func (fs *FileSystem) Remove(id interface{}) error {
	return fs.files().RemoveId(id)
}

func (fs *FileSystem) files() *mgo.GridFS {
	return fs.session.database.GridFS(fs.config.Name)
}
