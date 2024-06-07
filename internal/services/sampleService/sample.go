package sampleService

/*
	Discription: This is a sample service used to implement business logic.
*/

import (
	"goink/config"
	"time"

	"goink/pkg/utils/jwtUtil"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT() will use the signing method and secret key from the config to generate a JWT.
func GenerateJWT(ctx *gin.Context) (string, error) {
	type Example struct {
		UserID string `json:"userID"`
		jwt.RegisteredClaims
	}

	claims := Example{
		"ExampleID",
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(
				time.Duration(config.Middleware.Auth.JwtExpires) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	signedMethod := jwtUtil.JudgeJWTMethod(config.Middleware.Auth.JwtSigningMethod)
	var token *jwt.Token

	if config.Middleware.Auth.JwtSigningMethod[:2] == "HS" {
		token = jwt.NewWithClaims(signedMethod.(*jwt.SigningMethodHMAC), claims)
	} else if config.Middleware.Auth.JwtSigningMethod[:2] == "RS" {
		token = jwt.NewWithClaims(signedMethod.(*jwt.SigningMethodRSA), claims)
	} else if config.Middleware.Auth.JwtSigningMethod[:2] == "ES" {
		token = jwt.NewWithClaims(signedMethod.(*jwt.SigningMethodECDSA), claims)
	}
	tokenString, err := token.SignedString([]byte(config.Middleware.Auth.JwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
