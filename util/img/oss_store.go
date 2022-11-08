package img

import (
	"io"
)

type OSSImageStore struct {
}

func (s *OSSImageStore) SaveImage(file io.Reader, tpe ImageType, id int) {
	//TODO implement me
	panic("implement me")
}

func (s *OSSImageStore) GetImagePath(tpe ImageType, id int) string {
	//TODO implement me
	panic("implement me")
}
