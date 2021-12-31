package main

import (
	"io"
	"log"
	"os"

	"github.com/Fakorede/gin-app/controllers"
	"github.com/Fakorede/gin-app/docs" // swagger generated files
	"github.com/Fakorede/gin-app/middlewares"
	"github.com/Fakorede/gin-app/repository"
	"github.com/Fakorede/gin-app/router"
	"github.com/Fakorede/gin-app/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"     // swagger embedded files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var (
	videoRepository repository.VideoRepository

	loginService services.LoginService
	jwtService   services.JWTService
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

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Go Gin API"
	docs.SwaggerInfo.Description = "A RESTful API built with Golang, Gin framework and Gorm"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// defer videoRepository.CloseDB()

	setupLogOutput()
	setupDotEnv()
	setupDependencies()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	// server.Use(middlewares.BasicAuth())
	// server.Use(gindump.Dump())

	routes(server)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}

func routes(server *gin.Engine) {
	authAPIs := router.NewAuthAPIs(loginController)
	videoAPIs := router.NewVideoAPIs(videoController)

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", authAPIs.Authenticate)
		}

		videos := apiRoutes.Group("videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPIs.GetVideos)

			videos.POST("", videoAPIs.CreateVideo)

			videos.PUT(":id", videoAPIs.UpdateVideo)

			videos.DELETE(":id", videoAPIs.DeleteVideo)
		}
	}

	viewRoutes := server.Group("/views")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}
}
