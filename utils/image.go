package utils

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"log"

	"github.com/nfnt/resize"
)

func ResizeImg(img image.Image) image.Image{
	resizedImage := resize.Resize(250,250,img,resize.Lanczos2)
	return resizedImage
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func imgToBytes(img image.Image) []byte {
    var opt jpeg.Options
    opt.Quality = 80

    buff := bytes.NewBuffer(nil)
    err := jpeg.Encode(buff, img, &opt)
    if err != nil {
        log.Fatal(err)
    }

    return buff.Bytes()
}

func ConvertToBase64(img image.Image) string{
	imgByte:= imgToBytes(img)

	bs64string:= "data:image/jpeg;base64,"+toBase64(imgByte)
	
	return bs64string
}
