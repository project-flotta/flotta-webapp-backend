package devices

import (
	"github.com/ahmadateya/flotta-webapp-backend/pkg/s3"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Init(r *gin.Engine) {
	// use cors middleware to allow cross-origin requests
	r.Use(CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})

	// init device routes
	h := Handler{
		S3: s3.InitS3Client(),
	}
	r.GET("/devices", h.ListDevices)
	r.GET("/devices/:device", h.GetDeviceData)
	//r.GET("/devices/:device/network", h.GetNetworkData)
}

// CORSMiddleware returns a middleware handler that adds proper CORS headers to each request.
// this could be moved to a separate package if it is used in other places
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
