package main

import (
	"github.com/gin-gonic/gin"
)

/* นิยาม struct */
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func getUsers(c *gin.Context) {
	user := []User{{ID: "1", Name: "Pongrapee"}}

	c.JSON(200, user)
}

func main() {
	/* สร้างตัวเราท์ขึ้นมา (มาจาก gin) */
	r := gin.Default()

	r.GET("users", getUsers)

	r.Run(":8080")
}
