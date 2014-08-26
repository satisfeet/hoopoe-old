package mongo

import (
	"io"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type File interface {
	Id() interface{}
	SetId(interface{})

	io.ReadWriteCloser
}

type FileSystem struct {
	config  Config
	session *Session
}

func (fs *FileSystem) New(f File) error {
	var err error

	f, err = fs.files().Create("")
	f.SetId(bson.NewObjectId())

	return err
}

func (fs *FileSystem) Open(q *Query, f File) error {
	id, err := q.id()

	if err != nil {
		return err
	}

	f, err = fs.files().OpenId(id)

	return err
}

func (fs *FileSystem) Remove(q *Query) error {
	id, err := q.id()

	if err != nil {
		return err
	}

	return fs.files().RemoveId(id)
}

func (fs *FileSystem) files() *mgo.GridFS {
	return fs.session.database.GridFS(fs.config.Name)
}
