package middleware

/*
	Discription: You can write down your own auth middleware here.
*/

import (
	"fmt"
	"goink/config"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthBearer() is used to verify the JWT token
func AuthBearer() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")

		if tokenString == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if config.Middleware.Auth.JwtSigningMethod[:2] == "HS" {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
			} else if config.Middleware.Auth.JwtSigningMethod[:2] == "RS" {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
			} else if config.Middleware.Auth.JwtSigningMethod[:2] == "ES" {
				if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
			}

			return []byte(config.Middleware.Auth.JwtSecret), nil
		})

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		userID := claims["userID"].(string)
		ctx.Set("userID", userID)
		ctx.Next()
	}
}

// AuthCustom() is used for verifying a request with an valid X-Auth header
func AuthCustom() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authToken := ctx.Request.Header.Get("X-Auth")

		if authToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		if authToken != config.Middleware.Auth.XAuthToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			return
		}

		ctx.Next()
	}
}
