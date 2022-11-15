package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

type config struct {
	port int
}

type application struct {
	config config
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	app := &application{
		config: cfg,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	fmt.Printf("starting server on port %s", srv.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
