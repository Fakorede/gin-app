package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Fakorede/gin-app/controllers"
	"github.com/Fakorede/gin-app/middlewares"
	"github.com/Fakorede/gin-app/services"
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService services.VideoService = services.New()
	videoController controllers.VideoController = controllers.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	server.Use(middlewares.BasicAuth())
	server.Use(gindump.Dump())

	server.GET("/", func (ctx *gin.Context)  {
		ctx.JSON(200, gin.H{"message": "Hello"})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
	})

	server.Run(":8085")
}