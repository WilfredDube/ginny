package internal

import (
	"fmt"
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
func (app *Application) Routes() *gin.Engine {
	r := gin.Default()

	r.Use(app.recoverPanic())

	// TODO: Use app.handler.Pusher() -> HTTP/2 & https required
	r.Static("/assets", "ui/assets")
	r.LoadHTMLGlob("ui/html/*")

	r.Handle("GET", "/", app.handler.Home)
	r.Handle("GET", "/snippet", app.handler.ShowSnippet)
	r.Handle("GET", "/snippet/create", app.handler.PrepareSnippet)
	r.Handle("GET", "/snippet/create/:id", app.handler.PrepareSnippet)
	r.Handle("POST", "/snippet/create/:id", app.handler.CreateSnippet)

	return r
}

func (app *Application) recoverPanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.Header("Connection", "close")
				http.ServerError(ctx, fmt.Errorf("%s", err))
			}
		}()

		ctx.Next()
	}
}
