package routes

import "github.com/gin-gonic/gin"

func InitHomeRoutes(rg *gin.RouterGroup) {
	routerGroup := rg.Group("/")
	routerGroup.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"API VERSION": "1.0.0",
		})
	})
}
