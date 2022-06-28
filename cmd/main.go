package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WilfredDube/ginny/internal"
)

func main() {
	var cfg internal.Config

	flag.IntVar(&cfg.Port, "port", 4000, "HTTP network port")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	app := &internal.Application{
		Config: cfg,
		Logger: logger,
	}

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.Port),
		Handler:  app.Routes(),
		ErrorLog: errlog,
	}

	logger.Println("Starting server on", cfg.Port)
	err := server.ListenAndServe()
	errlog.Fatal(err)
}
