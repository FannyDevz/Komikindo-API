package main

import (
	"log"

	"baca-manga/helpers"
	routes "baca-manga/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := helpers.GetKeyEnv("PORT")
	if port == "" {
		port = "8080"
	}
	router := InitRoute()
	log.Fatal(router.Run(":" + port))
}

func InitRoute() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(cors.Default())
	RouteApi := router.Group("/api")
	routes.Routes(RouteApi)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})
	return router
}
