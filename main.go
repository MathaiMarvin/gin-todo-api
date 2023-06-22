package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
)

func main(){

	//We initialize a new gin Engine with the correct middleware.
	router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// })
	router.Run(":8080")
}