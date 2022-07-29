package main

import (
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/api"
	"github.com/ahmadateya/flotta-webapp-backend/config"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/s3"
	"log"
	"net/http"
	"time"
)

func main() {

	// read configurations from env file
	cfg, err := config.NewConfig("./config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file : %v", err)
	}

	s3Client := s3.InitS3Client()
	s3Client.ListTopLevelFolders()

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
