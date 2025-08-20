package main

import (
	"github.com/gin-gonic/gin"
)

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {
	r := gin.Default()

	r.GET("/users", func(c *gin.Context) {
		user := []user{{ID: "1", Name: "Pongrapee"}}
		c.JSON(200, user)
	})

	r.Run(":8080")
}
