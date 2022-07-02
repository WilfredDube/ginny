package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/WilfredDube/ginny/internal"
	"github.com/WilfredDube/ginny/internal/db"
)

func main() {
	var cfg internal.Config

	flag.IntVar(&cfg.Port, "port", 4000, "HTTP network port")
	flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
	flag.StringVar(&cfg.DSN, "db-dsn", os.Getenv("SNIPPET_DSN"), "PostgreSQL DSN")
	var migrationDir = flag.String("migration.files", "migrations", "Directory where the migration files are located ?")

	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	errlog := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	db, dbconn, err := openDB(cfg.DSN)
	if err != nil {
		errlog.Fatal(err)
	}

	defer func() {
		_ = db.Close()
		fmt.Println("database connection closed")
	}()

	driver, err := postgres.WithInstance(dbconn, &postgres.Config{})
	if err != nil {
		errlog.Fatal(err)
	}

	migrateDB(*migrationDir, driver)

	app := internal.NewApplication(cfg, db, logger, errlog)

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

func migrateDB(dir string, driver database.Driver) {
	source := fmt.Sprintf("file://%v", dir)

	m, err := migrate.NewWithDatabaseInstance(source, "postgres", driver)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("Database migrated: NO CHANGE", source)
		} else {
			log.Fatalf("An error occurred while syncing the database.. %v", err)
		}
		return
	}

	log.Println("Database migrated")
}
