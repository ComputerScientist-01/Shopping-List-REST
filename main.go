package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShoppingList struct {
	Name string `json:"name"`
}

func main() {

	shoppingLists := make([]ShoppingList, 0)

	router := gin.New()

	router.GET(
		"/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message":"pong",
			})
		},
	)

	router.POST(
		"/shopping-list/add",
		func(c *gin.Context) {
			var sl ShoppingList
			if err := c.BindJSON(&sl); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "bad request",
				})
				return
			}

			shoppingLists = append(shoppingLists, sl)
			c.JSON(http.StatusOK, gin.H{
				"message": "Added shopping list",
			})
		},
	)

	log.Println("starting server on port 9030")

	router.Run(":9030")
}
