package handlers

import (
	"github.com/ahmadateya/flotta-webapp-backend/pkg/s3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HelloServer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func ListMachines(c *gin.Context) {
	// get machine names from S3 top level folders
	client := s3.InitS3Client()
	machines := client.ListTopLevelFolders()

	// trim slash from machine names
	for i, machine := range machines {
		machines[i] = strings.TrimSuffix(machine, "/")
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": []map[string]interface{}{
			{
				"machines": machines,
			},
		},
	})
}
