package api

import (
	// "github.com/gin-contrib/cors"

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

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)

	protected := router.Group("/", middleware.JWTMiddleware())
	protected.GET("/profile/:id", h.Profile)

	return router
}
