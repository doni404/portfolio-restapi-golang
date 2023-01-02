package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/doni404/portfolio-restapi-golang/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Check environment variable (Load .env file)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("SITE_ENV")
	if env == "" {
		env = "DEV"
	}

	// DB_USER := os.Getenv("DB_USER")
	// DB_PASSWORD := os.Getenv("DB_PASSWORD")
	// DB_NAME := os.Getenv("DB_NAME")
	// DB_HOST := os.Getenv("DB_HOST")
	// DB_PORT := os.Getenv("DB_PORT")
	APP_PORT := os.Getenv("APP_PORT")

	// GORM DB Connection
	// Ask DB admin about the connection
	// dsn := DB_USER + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// if errDB != nil {
	// 	log.Fatal("DB Connection Error")
	// }

	/* Golang Structure */
	/* Main -> Handler -> Service -> Repository -> DB (MysQL) */

	// Start API using GIN & Handle CORS
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		errorMessage := gin.H{"error": "Page not found"}
		response := helper.APIResponse("PAGE_NOT_FOUND", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusNotFound, response)
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"}
	corsConfig.AllowOrigins = []string{"*"}
	router.Use(cors.New(corsConfig))

	// Versioning API
	api := router.Group("/v1")

	api.GET("/blogs", func(c *gin.Context) {
		response := helper.APIResponse("Successfully get all blogs", http.StatusOK, "success", "Good !")
		c.JSON(http.StatusOK, response)
	})

	// Send Contact Test
	// var contactRequest contact.ContactRequest
	// contactRequest.Name = "Doni"
	// contactRequest.Email = "d.putra@global-coding.com"
	// contactRequest.Phone = "087818231232"
	// contactRequest.Message = "Hello, Mail Test !"
	// helper.SendContactEmail(contactRequest, "support@doniputra.com")

	// Setup Custom Port by .env
	openPort := fmt.Sprint(":", APP_PORT)
	router.Run(openPort)
	fmt.Print("REST API run on port : ", APP_PORT)

}
