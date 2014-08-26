package store

import "github.com/satisfeet/hoopoe/store/common"

var imageName = "images"

type ImageStore struct {
	*common.Files
}

func NewImageStore(s *common.Session) *ImageStore {
	return &ImageStore{
		Files: common.NewFiles(common.Config{
			Name: imageName,
		}, s),
	}
}
