package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type individual struct {
	Name   string `json:"name"`
	Age    int8   `json:"age"`
	Weight int8   `json:"weight"`
}

var people = []individual{
	{Name: "Peter", Age: 22, Weight: 80},
	{Name: "Emily", Age: 25, Weight: 55},
	{Name: "Chris", Age: 27, Weight: 85},
}

func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, people)
}

func main() {
	router := gin.Default()
	router.GET("/people", getPeople)
	router.Run("localhost:8080")
}
