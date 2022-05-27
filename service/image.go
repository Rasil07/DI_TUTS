package service

import (
	"bytes"
	"dependency_injection_tut/utils"
	"image"
	"image/jpeg"
	"io/ioutil"
	"log"
	"mime/multipart"
)

type ImageService struct{}

func NewImageService() *ImageService{
	return &ImageService{}
}


func(ims *ImageService) Upload(files []*multipart.FileHeader){
	log.Println("Image service samma pugyo",files)
	
}

func(ims *ImageService) Resize(files []*multipart.FileHeader) ([]string,error){
	var imgs []image.Image
	var base64Strings []string

	for _,file :=range files{
		fl, err:= file.Open()
		if err!=nil{
			return base64Strings,err
		}
		defer fl.Close()

		if err!=nil{
			log.Println("Couldnot open file")
			return base64Strings,err

		}
		flRead,err:= ioutil.ReadAll(fl) 
		if err!=nil{
			log.Println("Couldnot open file")
			return base64Strings,err
		}
		m, err := jpeg.Decode(bytes.NewReader(flRead))
		if err !=nil{
			log.Println("Couldnot decode image")
			return base64Strings,err

		}
		resizedImage:= utils.ResizeImg(m)
		imgs = append(imgs,resizedImage)
	}
   
   for _,resizedImg:= range imgs{
	   base64Strings = append(base64Strings, utils.ConvertToBase64(resizedImg))
   }
	return base64Strings,nil
} 