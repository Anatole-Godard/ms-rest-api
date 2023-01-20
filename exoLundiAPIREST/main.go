package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type params struct {
	Param1 int `json:"param1"`
	Param2 int `json:"param2"`
}

func postMethod(c *gin.Context) {

	var requestBody params

	if err := c.BindJSON(&requestBody); err != nil {
		c.IndentedJSON(
			http.StatusBadRequest,
			"pas contente",
		)
	}

	fmt.Printf("%+v\n", requestBody)
	var addition = requestBody.Param1 + requestBody.Param2
	c.IndentedJSON(http.StatusOK, addition)
}

func main() {
	router := gin.Default()
	router.POST("/somme", postMethod)

	router.Run("localhost:9090")
}
