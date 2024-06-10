package middleware

import (
	"net/http"
	"strings"

	t "api-gateway/api/token"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		url := (ctx.Request.URL.Path)

		if strings.Contains(url, "swagger") || (url == "/v1/login") {
			ctx.Next()
			return
		} else if isValid, err := t.ValidateToken(token); !isValid && err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
