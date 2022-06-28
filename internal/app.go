package internal

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	controller "github.com/WilfredDube/ginny/internal/http"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Config   Config
	Logger   *log.Logger
	ErrorLog *log.Logger
	handler  controller.SnippetHandler
}

func (app *Application) Serve() error {
	// idleConnsClosed := make(chan struct{})

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", app.Config.Port),
		Handler:  app.Routes(),
		ErrorLog: app.ErrorLog,
	}

	shutdownErr := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		app.ErrorLog.Println("caught signal", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		shutdownErr <- server.Shutdown(ctx)
	}()

	app.Logger.Println("starting server")

	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErr
	if err != nil {
		return err
	}

	app.Logger.Println("stopping server", map[string]string{
		"addr": server.Addr,
	})

	return nil
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
				controller.ServerError(ctx, fmt.Errorf("%s", err))
			}
		}()

		ctx.Next()
	}
}
