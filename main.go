package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.GET(
		"/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":"pong",
			})
		},
	)

	log.Println("starting server on port 9030")

	router.Run(":9030")
}
