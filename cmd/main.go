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

	db, err := openDB(cfg.DSN)
	if err != nil {
		errlog.Fatal(err)
	}

	defer func() {
		_ = db.Close()
		fmt.Println("database connection closed")
	}()

	app := internal.NewApplication(cfg, logger, db)

	err = app.Serve()
	if err != nil {
		errlog.Fatal(err)
	}
}

func openDB(dsn string) (*db.Queries, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect%w", err)
	}

	db := db.New(conn)

	return db, nil
}
