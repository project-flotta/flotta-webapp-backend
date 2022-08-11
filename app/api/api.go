package api

import (
	"github.com/ahmadateya/flotta-webapp-backend/api/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	// use cors middleware to allow cross-origin requests
	r.Use(CORSMiddleware())

	// Initialize the App Routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	r.GET("/devices", handlers.ListDevices)
	r.GET("/devices/:device", handlers.GetNetworkTopologyData)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept-Language,  Authorization, accept, origin, Cache-Control, X-Requested-With, Session-Key, session-key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT , DELETE , PATCH")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
