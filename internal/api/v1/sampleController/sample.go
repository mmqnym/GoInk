package sampleController

/*
	Discription: This is a sample controller used to provide user interaction routing and to communicate
	with business logic.
*/

import (
	"goink/internal/models/responseModel"
	"goink/internal/services/sampleService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags API
// @Summary Test the server is running normally
// @Description This API will return 200 if the server is running normally
// @ID v1-sample-ping
// @Produce application/json
//
// @Success 200 {object} responseModel.Ping "Success" response=`{"message": "pong"}`
// @Router /sample/ping [get]
func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, responseModel.Ping{
		Message: "pong",
	})
}

// @Tags API
// @Summary Generating JWTs for Authentication
// @Description This API will generate JWT which users can use on some APIs that require authentication
// @ID v1-sample-generate-jwt
// @Produce application/json
//
// @Success 200 {object} responseModel.GenerateJWT "Success" response=`{"token": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJFeGFtcGxlSWQiLCJleHAiOjE3MTc3NDYyMTUsIm5iZiI6MTcxNzY1OTgxNSwiaWF0IjoxNzE3NjU5ODE1fQ.BOYdOVsiPbK5nfhsYzX4G_EciVlbZ3LM2pCUQFjRwZo"}`
// @Failure 500 {object} responseModel.Error "Failure" response=`{"error": "something went wrong"}`
// @Router /sample/jwt/get [get]
func GenerateJWT(ctx *gin.Context) {

	tokenString, err := sampleService.GenerateJWT(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responseModel.Error{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responseModel.GenerateJWT{
		Token: "Bearer " + tokenString,
	})
}

// @Tags API
// @Summary Verifying JWTs
// @Description This API will check if the JWT is valid
// @ID v1-sample-test-jwt
// @Produce application/json
//
// @Param Authorization header string true "Bearer"
// @Success 200 {object} responseModel.TestJWT "Success" response=`{"userID": "ExampleID"}`
// @Router /sample/jwt/test [post]
func TestJWT(ctx *gin.Context) {
	id := ctx.GetString("userID")
	ctx.JSON(http.StatusOK, responseModel.TestJWT{
		UserID: id,
	})
}
