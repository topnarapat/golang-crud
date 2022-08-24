package main

import (
	"os"

	"example.com/gin-backend-api/configs"
	"example.com/gin-backend-api/routes"
	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := SetupRouter()
	router.Run(":" + os.Getenv("GO_PORT"))
}

func SetupRouter() *gin.Engine {
	// Load .env
	godotenv.Load(".env")

	gin.SetMode(os.Getenv("GIN_MODE"))

	// connect db (postgres)
	configs.Connection()

	// conct db (mongo db)
	configs.MongoConnection()

	router := gin.Default()

	// limit request size
	var maxBytes int64 = 1024 * 1024 * 10 // 10MB
	router.Use(limits.RequestSizeLimiter(maxBytes))
	// Serving static files
	router.Static("/api/v1/public/images/", "./public/images/")

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		// MaxAge: 12 * time.Hour,
	}))

	apiV1 := router.Group("/api/v1") //localhost:3000/api/v1

	routes.InitHomeRoutes(apiV1)
	routes.InitUserRoutes(apiV1)
	routes.InitLogRoutes(apiV1)
	routes.InitUploadRoutes(apiV1)

	return router
}
