package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.1.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	cfg.env = "dev"
	cfg.port = 4000
	//flag.IntVar(&cfg.port, "port", 4000, "API server port")
	//flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	//flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	logger.Println(fmt.Sprintf("Server is running on port %s", srv.Addr))
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
