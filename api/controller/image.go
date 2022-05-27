package controller

import (
	"dependency_injection_tut/service"
	"net/http"

	"github.com/gin-gonic/gin"
)


type ImageController struct{
	imageService *service.ImageService
}

func NewImageController(ser *service.ImageService) *ImageController{
	return &ImageController{imageService: ser}
}

func (imc *ImageController) HandleResize(c *gin.Context){
	form,_ := c.MultipartForm()
	files := form.File["file"]
   resiezedImgs,err := imc.imageService.Resize(files)
   
   if err!=nil{
	   c.JSON(http.StatusBadRequest,gin.H{
		   "error":err.Error(),
	   })
	   return
	}

   c.JSON(200,gin.H{
	   "img":resiezedImgs,
   })
}


func (imc *ImageController) HandleUpload(c *gin.Context){}