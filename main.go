package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		// health check
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = ":8081"
	}
	err := router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
