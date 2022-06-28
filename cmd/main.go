package main

import (
	"flag"
	"log"
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
		Config:   cfg,
		Logger:   logger,
		ErrorLog: errlog,
	}

	err := app.Serve()
	if err != nil {
		errlog.Fatal(err)
	}
}
