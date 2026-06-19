package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(RequestIDMiddleware())
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	r.GET("/read-me-at-browser", func(c *gin.Context) {
		requestId, _ := c.Get("request_id")
		log.Println("Request Id:", requestId)
		c.File("files/OperatingSystemsThreeEasyPieces.pdf")
	})

	r.GET("download", func(c *gin.Context) {
		requestId, _ := c.Get("request_id")
		log.Println("Request Id:", requestId)
		c.FileAttachment("files/OperatingSystemsThreeEasyPieces.pdf", "OperatingSystemsThreeEasyPieces.pdf")
	})

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	r.Run(port)
}
