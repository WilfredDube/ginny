package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Home diplays the home page containing a list of all snippets.
func Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello from home...")
}

// ShowSnippet displays a single selected snippet.
func ShowSnippet(c *gin.Context) {
	c.String(http.StatusOK, "Display a specific snippet...")
}

// PrepareSnippet prepares an URL with a UUID for the creation of a snippets. Makes the
// creation idempotent.
func PrepareSnippet(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		// Generate GUID to make call idempotent
		location := c.FullPath() + "/" + uuid.NewString()

		c.Redirect(http.StatusMovedPermanently, location)
	}

	c.String(http.StatusOK, "Create new snippet..."+id)
}

// CreateSnippet creates a snippet. Call is idempotent.
func CreateSnippet(c *gin.Context) {
	id := c.Param("id")

	c.String(http.StatusOK, "Create new snippet..."+id)
}

// Creates a router
func setUpRoutes() *gin.Engine {
	r := gin.Default()

	r.Handle("GET", "/", Home)
	r.Handle("GET", "/snippet", ShowSnippet)
	r.Handle("GET", "/snippet/create", PrepareSnippet)
	r.Handle("GET", "/snippet/create/:id", PrepareSnippet)
	r.Handle("POST", "/snippet/create/:id", CreateSnippet)

	return r
}

func main() {
	router := setUpRoutes()
	router.Run()
}
