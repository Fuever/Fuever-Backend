package img

import (
	"io"
)

type LocalImageStore struct {
}

func (s *LocalImageStore) SaveImage(file io.Reader, tpe ImageType, id int) {
	//TODO implement me
	panic("implement me")
}

func (s *LocalImageStore) GetImagePath(tpe ImageType, id int) string {
	//TODO implement me
	panic("implement me")
}
