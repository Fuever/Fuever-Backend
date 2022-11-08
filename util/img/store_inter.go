package img

import (
	"io"
)

type ImageType = int

const (
	USER_AVATAR ImageType = iota
	NEW_COVER
	ANNIVERSARY_COVER
)

type ImageStoreManager interface {
	// SaveImage 保存图片
	//
	// 值得一提的是
	// 要根据tpe决定路径
	// 反正存储就是有路径依赖
	SaveImage(file io.Reader, tpe ImageType, id int)
	// GetImagePath 获取图片的存储路径
	//
	// 考虑到后续可能会使用阿里云OSS
	// 这边返回的可能是本地路径
	// 也可能是一个http地址
	GetImagePath(tpe ImageType, id int) string
}

func main() {

}
