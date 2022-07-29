package main

import (
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/api"
	"github.com/ahmadateya/flotta-webapp-backend/config"
	"log"
	"net/http"
	"time"
)

func main() {

	// read configurations from env file
	cfg, _ := config.NewConfig("./config.yaml")

	// Start the server
	r := api.Init()

	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
