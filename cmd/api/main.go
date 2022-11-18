package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"net/http"
)

//go:embed version.txt
var version string

type config struct {
	port int
	env  string
}

type application struct {
	config config
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|stating|production)")

	flag.Parse()
	app := &application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	fmt.Printf("starting %s server on port %s", cfg.env, srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
