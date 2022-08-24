package routes

import (
	uploadcontroller "example.com/gin-backend-api/controllers/upload"
	"github.com/gin-gonic/gin"
)

func InitUploadRoutes(rg *gin.RouterGroup) {

	routerGroup := rg.Group("/upload")

	//{{domain_url}}/api/v1/upload
	routerGroup.POST("/", uploadcontroller.UploadImage)

}
