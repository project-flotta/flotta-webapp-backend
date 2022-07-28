package main

import (
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/api"
	"github.com/ahmadateya/flotta-webapp-backend/helpers"
	"log"
	"net/http"
	"time"
)

func main() {

	// read configurations from env file
	c, _ := helpers.NewConfig("./config.yaml")

	// Start the server
	r := api.Init()
	srv := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf("%s:%s", c.Server.Host, c.Server.Port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
