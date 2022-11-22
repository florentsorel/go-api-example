package main

import (
	"context"
	"database/sql"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rtransat/go-api-example/internal/data"
	"github.com/rtransat/go-api-example/internal/jsonlog"

	_ "github.com/go-sql-driver/mysql"
)

//go:embed version.txt
var version string

type config struct {
	port int
	env  string
	db   struct {
		dsn             string
		maxLifetimeConn string
		maxOpenConns    int
		maxIdleConns    int
	}
}

type application struct {
	config config
	logger *jsonlog.Logger
	models data.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|stating|production)")

	dsn := os.Getenv("API_DB_DSN")
	if len(dsn) == 0 {
		dsn = "root:toor@/golang_api?parseTime=true"
	}

	flag.StringVar(&cfg.db.dsn, "db-dsn", dsn, "MySQL DSN")
	flag.StringVar(&cfg.db.maxLifetimeConn, "db-max-lifetime-conn", "3m", "MySQL max lifetime connection")
	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 10, "MySQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 10, "MySQL max idle connections")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := openDB(cfg)
	if err != nil {
		logger.PrintFatal(err, nil)
	}
	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:     fmt.Sprintf(":%d", cfg.port),
		Handler:  app.routes(),
		ErrorLog: log.New(logger, "", 0),
	}

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  cfg.env,
	})

	err = srv.ListenAndServe()
	if err != nil {
		logger.PrintFatal(err, nil)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	duration, err := time.ParseDuration(cfg.db.maxLifetimeConn)
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(duration)

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
