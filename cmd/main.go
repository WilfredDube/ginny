package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Home diplays the home page containing a list of all snippets.
func Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello from home...")
}

// ShowSnippet displays a single selected snippet.
func ShowSnippet(c *gin.Context) {
	c.String(http.StatusOK, "Display a specific snippet...")
}

// CreateSnippet creates a snippet. The creation is idempotent.
func CreateSnippet(c *gin.Context) {
	c.String(http.StatusOK, "Create new snippet...")
}

// Creates a router
func setUpRoutes() *gin.Engine {
	r := gin.Default()

	r.Handle("GET", "/", Home)
	r.Handle("GET", "/snippet", ShowSnippet)
	r.Handle("GET", "/snippet/create", CreateSnippet)

	return r
}

func main() {
	router := setUpRoutes()
	router.Run()
}
