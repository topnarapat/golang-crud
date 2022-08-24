package routes

import (
	logcontroller "example.com/gin-backend-api/controllers/log"
	"github.com/gin-gonic/gin"
)

func InitLogRoutes(rg *gin.RouterGroup) {

	routerGroup := rg.Group("/logs")

	//{{domain_url}}/api/v1/logs
	routerGroup.GET("/", logcontroller.GetLog)

	//{{domain_url}}/api/v1/logs/new
	routerGroup.POST("/new", logcontroller.InsertLog)

}
