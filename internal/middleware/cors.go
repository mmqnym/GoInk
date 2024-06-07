package middleware

/*
	Discription: You can write down your own cors middleware here.
*/

import (
	"goink/config"
	"goink/pkg/logger"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		origin := ctx.Request.Header.Get("Origin")

		if ok := isAllowedOrigin(origin); ok {
			ctx.Header("Access-Control-Allow-Origin", origin)
			ctx.Header("Access-Control-Allow-Methods", config.Middleware.Cors.AllowMethods)
			ctx.Header("Access-Control-Allow-Headers", config.Middleware.Cors.AllowHeaders)
			ctx.Header("Access-Control-Max-Age", config.Middleware.Cors.MaxAge)
			ctx.Header("Access-Control-Allow-Credentials", config.Middleware.Cors.AllowCredentials)
		}

		defer func() {
			if err := recover(); err != nil {
				logger.Sugar.Error(err.(error).Error())
			}
		}()

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}

func isAllowedOrigin(origin string) bool {
	if config.Middleware.Cors.AllowOrigins[0] == "*" {
		return true
	}

	for _, allowOrigin := range config.Middleware.Cors.AllowOrigins {
		if allowOrigin == origin {
			return true
		}
	}

	return false
}
