package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func homePage(c *gin.Context) {

	c.JSON(200, gin.H{

		"message": "Hi! I am from Home Page",
	})
}
func getdetails(c *gin.Context) {

	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{

		"name": name,
		"age":  age,
	})
}
func main() {
	fmt.Println("Hi! I am Ritwik")

	r := gin.Default()
	r.GET("/postq", getdetails)
	r.GET("/", homePage)
	r.Run()

}
