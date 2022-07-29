package api

import (
	"github.com/ahmadateya/flotta-webapp-backend/api/handlers"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	r.GET("/", handlers.HelloServer)
	r.GET("/machines", handlers.ListMachines)
}
