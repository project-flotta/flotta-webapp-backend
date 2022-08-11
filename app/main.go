package main

import (
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/api"
	"github.com/ahmadateya/flotta-webapp-backend/config"
	"github.com/gin-gonic/gin"
)

func main() {

	// read configurations from env file
	cfg, err := config.NewConfig("./config.yaml")
	if err != nil {
		fmt.Printf("Error reading config file : %v", err)
	}

	// Start the server
	router := gin.New()
	api.Init(router)
	err = router.Run(fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port))
	if err != nil {
		fmt.Printf("Error Starting the server %v\n", err.Error())
	}
}
