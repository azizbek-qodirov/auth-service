package api

import (
	// "github.com/gin-contrib/cors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "auth-service/api/docs"
	"auth-service/api/handlers"
	"auth-service/api/middleware"
)

// @Title N10 group swagger UI
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(userHandler *handlers.HTTPHandler) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:7070"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	router.Use(middleware.Middleware())
	api := router.Group("/v1")
	api.POST("/register/user", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/user", userHandler.GetUser)
	api.GET("/users", userHandler.GetUsers)
	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return router
}
