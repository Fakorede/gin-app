package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Fakorede/gin-app/controllers"
	"github.com/Fakorede/gin-app/middlewares"
	"github.com/Fakorede/gin-app/repository"
	"github.com/Fakorede/gin-app/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	videoRepository repository.VideoRepository
	loginService services.LoginService
	jwtService services.JWTService 
	videoService services.VideoService

	loginController controllers.LoginController 
	videoController controllers.VideoController
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func setupDotEnv() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

func setupDependencies() {
	videoRepository = repository.NewVideoRepository()

	loginService = services.NewLoginService()
	jwtService = services.NewJWTService()
	videoService = services.NewVideoService(videoRepository)

	loginController = controllers.NewLoginController(loginService, jwtService)
	videoController = controllers.NewVideoController(videoService)
}

func main() {
	// defer videoRepository.CloseDB()

	setupLogOutput()
	setupDotEnv()
	setupDependencies()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	server.Use(middlewares.BasicAuth())
	// server.Use(gindump.Dump())

	routes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	server.Run(":" + port)
}

func routes(server *gin.Engine) {
	// public routes
	server.GET("/", func (ctx *gin.Context)  {
		ctx.JSON(200, gin.H{"message": "Hello"})
	})

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)

		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		}
	})

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// authorized routes
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, videoController.FindAll())
		})
	
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
	
			ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
	
			ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
	
			ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
		})
	}
}