package main

import (
	"ecommerce/controllers"
	"ecommerce/database"
	"ecommerce/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	loadDatabase()
	r := gin.Default()

	r.GET("/books", controllers.GetAllBrands)
	r.GET("/books/:id", controllers.FindById)
	r.POST("/books", controllers.CreateBrand)
	r.DELETE("/books/:id", controllers.DeleteBrand)

	r.Run()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Brand{})
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
