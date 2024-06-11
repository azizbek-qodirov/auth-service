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

func NewRouter(h *handlers.HTTPHandler) *gin.Engine {
	router := gin.Default()

	url := ginSwagger.URL("swagger/doc.json")
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:7070"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	router.Use(cors.New(corsConfig))

	router.Use(middleware.Middleware())
	api := router.Group("/v1")
	api.POST("/register/user", h.Register)
	api.POST("/login", h.Login)
	api.GET("/user", h.GetUser)
	api.GET("/users", h.GetUsers)

	return router
}
