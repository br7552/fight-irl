package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/br7552/router"
)

type config struct {
	port       int
	mapsAPIkey string
}

type application struct {
	cfg config
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.mapsAPIkey, "mapsAPIkey", "",
		"Google Maps API key")
	flag.Parse()

	app := &application{
		cfg: cfg,
	}

	router := router.New()
	router.HandleFunc(http.MethodGet, "/", app.addrInfoHandler)
	router.HandleFunc(http.MethodGet, "/ip/:ip", app.meetingHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.cfg.port),
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}
