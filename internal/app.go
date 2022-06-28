package internal

import (
	"log"

	"github.com/WilfredDube/ginny/internal/http"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Config  Config
	Logger  *log.Logger
	handler http.SnippetHandler
}

// Creates a router
func Routes() *gin.Engine {
	r := gin.Default()

	r.Handle("GET", "/", http.Home)
	r.Handle("GET", "/snippet", http.ShowSnippet)
	r.Handle("GET", "/snippet/create", http.PrepareSnippet)
	r.Handle("GET", "/snippet/create/:id", http.PrepareSnippet)
	r.Handle("POST", "/snippet/create/:id", http.CreateSnippet)

	return r
}
