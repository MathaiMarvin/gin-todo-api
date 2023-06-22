package main

import (
	//"net/http"
	"github.com/gin-gonic/gin"
	"github.com/MathaiMarvin/gin-todo-api/routes"
)

func main(){

	//We initialize a new gin Engine with the correct middleware.
	router := gin.Default()
	// Importing the routes to gain access to the different routes.
	routes.UserRoute(router)
	router.Run(":8080")
}