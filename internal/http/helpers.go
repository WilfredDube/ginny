package http

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

// The serverError helper writes an error message ans stack trace to the errorlog
// then sends a generic 500 Internal server Error response to the user
func ServerError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	log.Println(trace)

	http.Error(c.Writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper sends a specific status code and corresponding description
// to the user.
func clientError(c *gin.Context, status int) {
	http.Error(c.Writer, http.StatusText(status), status)
}

//
func notFound(c *gin.Context) {
	clientError(c, http.StatusNotFound)
}
