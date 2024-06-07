package main

import (
	"goink/config"
	"goink/internal/api/v1/sampleController"
	"goink/internal/middleware"
	"os"
	"os/signal"

	"goink/pkg/logger"

	"github.com/gin-gonic/gin"

	_ "goink/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title goink
// @version 1.0
// @description GoInk is a template for quickly building a Go web server
// @contact.name 0xmmq
// @contact.url https://github.com/mmqnym
// @contact.email mail@mmq.dev
//
// @license.name MIT
// @license.url https://github.com/GoInk/blob/main/LICENSE
// @host localhost:9000
// @BasePath /api/v1
// @Scheme http
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go gracefulShutdown(c)

	if config.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	server := gin.Default()
	server.UseH2C = config.Server.Http2

	if config.Middleware.Cors.Enabled {
		server.Use(middleware.Cors())
	}

	regist(server)

	logger.Sugar.Infof("Server is running on port %s", config.Server.Port)
	server.Run(config.Server.Port)
}

// regist() will register the routes for version control
func regist(server *gin.Engine) {
	registV1(server)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// registV1() will register the v1 routes
func registV1(server *gin.Engine) {
	group := server.Group("/api/v1")

	// Add your routes here.

	sampleGroup := group.Group("/sample")
	sampleGroup.GET("/ping", sampleController.Ping)
	sampleGroup.GET("/jwt/get", sampleController.GenerateJWT)
	sampleGroup.POST("/jwt/test", middleware.AuthBearer(), sampleController.TestJWT)
}

// gracefulShutdown() will handle the shutdown signal for graceful shutdown
func gracefulShutdown(sig chan os.Signal) {
	<-sig

	logger.Sugar.Info("Server is shutting down...")
	logger.Logger.Sync()
	os.Exit(0)
}
