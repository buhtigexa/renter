package main

import (
	"database/sql"
	"flag"
	"github.com/buhtigexa/renter/pkg/models/postgresql"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDir string
	dsn       string
}

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	models   *postgresql.ResourceModel
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "user=web password=pass dbname=snippetbox sslmode=disable", "PostgreSQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &Application{errorLog: errorLog, infoLog: infoLog, models: &postgresql.ResourceModel{DB: db}}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn) // Change "mysql" to "postgres"
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
