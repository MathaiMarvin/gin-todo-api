package controller

// A controller function handles the logic for a specific route or endpoint in a web app.
// The controller function is the second argument to the router object's HTTP verb methods.
//When a request is made to that route, the router will invoke the corresponding controller function to handle the request. Helps in the separation of concerns routing and request handling making code modular and maintainable
import "github.com/gin-gonic/gin"

func UserController(c*gin.Context){

	c.String(200, "Hello World!")
}