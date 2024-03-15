package main

import (
	"ecommerce/database"
	"ecommerce/middleware"
	routes "ecommerce/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func CORSMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Credentials", "true")
		c.Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With",
		)
		c.Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(204)
		}

		return c.Next()
	}
}

func main() {
	loadEnv()
	loadDatabase()

	app := fiber.New()

	app.Use(CORSMiddleware())
	app.Use(middleware.AuthMiddleware)

	routes.RegisterRoutes(app)

	app.ListenTLS(":8443", "./cert.cert", "key.key")
	/**/ log.Fatal(app.Listen("0.0.0.0:8443"))
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
