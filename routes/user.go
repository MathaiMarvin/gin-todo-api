package routes

import (
	"github.com/MathaiMarvin/gin-todo-api/controller"
	"github.com/gin-gonic/gin"
)

//The function below takes a pointer to a gin.Engine object named router as its argument
func UserRoute(router *gin.Engine){

	//We use the router object to define the route for the /user endpoint
	// The second argument after the route is an anonymous function that takes a gin.Context object as its parameter. The function will be executed when a request is made to the endpoint.
	// router.GET("/", func(c *gin.Context) {
	// 	//c is the name of the parameter in the above and *gin.Context is the type of the parameter showcasing that c is a pointer to an object of type gin.Context. - THis type represents the context of an HTTP request and provides access to various request related information and methods
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World!",
	// 	})
	// 	//The gin.H is  a helper type provided by the gin framework to simplify working with maps in context of JSON responses.
	// })

	router.GET("/", controller.UserController)
}