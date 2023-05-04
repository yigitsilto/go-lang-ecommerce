package main

import (
	"ecommerce/controllers"
	"ecommerce/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	r := gin.Default()

	// Enable CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	r.Use(cors.New(config))

	r.GET("/brands", controllers.GetAllBrands)
	r.GET("/brands/:id", controllers.FindById)

	r.Run()
}

func loadDatabase() {
	database.Connect()
	//database.Database.AutoMigrate(&model.Brand{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
