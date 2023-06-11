package main

import (
	"crypto/tls"
	"ecommerce/database"
	routes "ecommerce/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
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

	routes.RegisterRoutes(r)

	// Sertifika ve anahtar dosyasının yolu
	certFile := "./cert.crt"
	keyFile := "./key.key"

	// TLS yapılandırmasını ayarla
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
	}

	// Sunucu oluştur ve çalıştır
	server := &http.Server{
		Addr:      ":8081",
		Handler:   r,
		TLSConfig: tlsConfig,
	}

	// HTTPS üzerinde 8080 portunda sunucuyu başlat
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
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
