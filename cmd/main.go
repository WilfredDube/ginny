package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/WilfredDube/ginny/internal"
	"github.com/WilfredDube/ginny/internal/db"
)

func main() {
	var cfg internal.Config

	flag.IntVar(&cfg.Port, "port", 4000, "HTTP network port")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.DSN, "db-dsn", os.Getenv("SNIPPET_DSN"), "PostgreSQL DSN")
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

func openDB(dsn string) (*db.Queries, *sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, nil, fmt.Errorf("pgx.Connect%w", err)
	}

	if err = conn.Ping(); err != nil {
		return nil, nil, err
	}

	db := db.New(conn)

	return db, conn, nil
}
}
