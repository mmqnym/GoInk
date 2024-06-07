package jwtUtil

import (
	"github.com/golang-jwt/jwt/v5"
)

// JudgeJWTMethod() is used to judge the jwt signing method
func JudgeJWTMethod(method string) interface{} {
	if method == "HS256" {
		return jwt.SigningMethodHS256
	} else if method == "HS384" {
		return jwt.SigningMethodHS384
	} else if method == "HS512" {
		return jwt.SigningMethodHS512
	} else if method == "RS256" {
		return jwt.SigningMethodRS256
	} else if method == "RS384" {
		return jwt.SigningMethodRS384
	} else if method == "RS512" {
		return jwt.SigningMethodRS512
	} else if method == "ES256" {
		return jwt.SigningMethodES256
	} else if method == "ES384" {
		return jwt.SigningMethodES384
	} else if method == "ES512" {
		return jwt.SigningMethodES512
	} else {
		return nil
	}
}
