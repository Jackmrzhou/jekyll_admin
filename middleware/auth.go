package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

var TrustTokens = map[string]bool{}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 4001,
				"message": "unauthorized",
			})
			ctx.Abort()
			return
		} else if _, ok := TrustTokens[token]; !ok {
			logrus.Warn("invalid token caught")
			ctx.JSON(http.StatusOK, gin.H{
				"code": 4001,
				"message": "invalid token",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}