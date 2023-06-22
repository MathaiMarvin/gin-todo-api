package routes

import "github.com/gin-gonic/gin"


//The function below takes a pointer to a gin.Engine object named router as its argument
func UserRoute(router *gin.Engine){

	//We use the router object to define the route for the /user endpoint
	// The second argument after the route is an anonymous function that takes a gin.Context object as its parameter. The function will be executed when a request is made to the endpoint.
	router.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
}