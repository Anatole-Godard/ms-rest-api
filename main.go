package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type params struct {
	param1 string
	param2 string
}

func postMethod(c *gin.Context) {
	var param params

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&param); err != nil {
		return
	}

	// Add the new album to the slice.
	fmt.Printf("%+v\n", param)
	c.IndentedJSON(http.StatusOK, param)
}

func main() {
	router := gin.Default()
	router.POST("/somme", postMethod)

	router.Run("localhost:9090")
}
