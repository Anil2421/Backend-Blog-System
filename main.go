package main

import (
	"log"

	"simpleblog/controllers"
	"simpleblog/initializers"
	"simpleblog/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	BlogController      controllers.BlogController
	BlogRouteController routes.BlogRouteController
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	BlogController = controllers.NewBlogController(initializers.DB)
	BlogRouteController = routes.NewRouteBlogController(BlogController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	BlogRouteController.BlogRoute(server)
	log.Fatal(server.Run(":" + config.ServerPort))

}
