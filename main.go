package main

import (
	"file_uploader/config"
	"file_uploader/db"
	"file_uploader/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.SetupCloudinary()

	dbErr := db.ConnectToDB()
	if dbErr != nil {
		log.Fatal("Error connecting to database:", dbErr)
	}
}

//documentation at localhost:3000/docs/index.html


// @title File uploader API
// @host localhost:3000
func main() {
	startGin()
}

func startGin() {
	gin.SetMode(gin.DebugMode)

	router := gin.Default()

	// Serve Swagger documentation
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.LoadFileRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("Error starting server: %s", err)
	}
}
