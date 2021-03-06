package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type config struct {
	port       int
	mapsAPIkey string
}

type application struct {
	cfg    config
	logger *log.Logger
}

func main() {
	var cfg config

	cfg.mapsAPIkey = os.Getenv("MAPKEY")

	flag.IntVar(&cfg.port, "port", 8080, "API server port")
	flag.Parse()

	logger := log.New(os.Stdout, "fight-irl-api", log.LstdFlags)

	app := &application{
		cfg:    cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.cfg.port),
		Handler:      app.routes(),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
