package handlers

import (
	"encoding/json"
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

func GetNetworkTopologyData(c *gin.Context) {
	machine := c.Param("machine")

	// get network topology data from S3
	client := s3.InitS3Client()
	networkTopologyFilename := client.GetMostRecentObjectNameInFolder(machine)

	// if the machine does not have any network data yet, return error
	if networkTopologyFilename == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": []map[string]interface{}{
				{
					"status": http.StatusNotFound,
					"title":  "no network topology data found",
					"detail": "no network topology data found for machine " + machine,
				},
			},
		})
		return
	}

	// read contents of network topology file from S3
	objContent := client.ReadObject(networkTopologyFilename)

	// parse objContent data to JSON
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(objContent), &jsonMap)

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": jsonMap,
	})
}
