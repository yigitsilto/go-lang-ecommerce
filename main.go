package main

import (
	"ecommerce/database"
	"ecommerce/middleware"
	routes "ecommerce/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	loadEnv()
	loadDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())
	r.Use(middleware.AuthMiddleware)

	routes.RegisterRoutes(r)

	r.RunTLS(":8443", "./cert.cert", "key.key")
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
