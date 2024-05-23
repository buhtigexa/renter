package main

import (
	"database/sql"
	"flag"
	"github.com/buhtigexa/renter/pkg/models/mysql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Addr      string
	StaticDic string
	dsn       string
}

type Application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	models   *mysql.ResourceModel
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.dsn, "dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(cfg.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	app := &Application{errorLog: infoLog, infoLog: errorLog, models: &mysql.ResourceModel{DB: db}}

	srv := &http.Server{
		Addr:     cfg.Addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.Addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
