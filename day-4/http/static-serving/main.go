package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Define a struct for User data
type User struct {
	Name string
	Age  int
}

func main() {
	// Create a default Gin router
	r := gin.Default()

	// Serve files (CSS for this case) from the "views/css" directory when "/css" URL is visited
	r.Static("/css", "views/css")

	// Load all HTML templates from "views/html" directory
	r.LoadHTMLGlob("views/html/*")

	// Handle HTTP GET request at route "/"
	r.GET("/", func(c *gin.Context) {

		// Prepare some user data
		users := []User{
			{Name: "John", Age: 22},
			{Name: "Jane", Age: 23},
		}

		// Render the HTML template named "index.gohtml" (located in "views/html directory")
		// Send the users data as the second parameter, which the template can use
		c.HTML(http.StatusOK, "index.gohtml", users)
	})

	// Run the server on port 8080 (default port)
	r.Run()
}
