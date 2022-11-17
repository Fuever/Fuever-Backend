package service

import (
	"fmt"
	"image/color"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func MakeCaptcha() (id, b64s string, err error) {
	var driver base64Captcha.Driver

	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor:         &color.RGBA{R: 3, G: 102, B: 214, A: 125},
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = captcha.Generate()
	return id, b64s, err
}

func VerifyCaptcha(id string, VerifyValue string) bool {
	fmt.Println(id, VerifyValue)
	if store.Verify(id, VerifyValue, true) {
		return true
	} else {
		return false
	}
}
