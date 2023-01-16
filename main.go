package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type params struct {
	param1 int
	param2 int
}

func postMethod(c *gin.Context) {
	var param params

	// Call BindJSON to bind the received JSON
	if err := c.BindJSON(&param); err != nil {
		return
	}

	// Add the new album to the slice.
	fmt.Printf("%+v\n", param)
	var addition = param.param1 + param.param2
	c.IndentedJSON(http.StatusOK, addition)
}

func main() {
	router := gin.Default()
	router.POST("/somme", postMethod)

	router.Run("localhost:9090")
}
