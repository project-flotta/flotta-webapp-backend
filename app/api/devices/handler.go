package devices

import (
	"encoding/json"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/s3"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Handler struct {
	S3 s3.S3
}

func (h *Handler) ListDevices(c *gin.Context) {
	// get machine names from S3 top level folders
	devices := h.S3.ListTopLevelFolders()

	// trim slash from machine names
	for i, device := range devices {
		devices[i] = strings.TrimSuffix(device, "/")
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": devices,
	})
}

func (h *Handler) GetNetworkTopologyData(c *gin.Context) {
	device := c.Param("device")

	// get network topology data from S3
	client := s3.InitS3Client()
	networkTopologyFilename := h.S3.GetMostRecentObjectNameInFolder(device)

	// if the machine does not have any network data yet, return error
	if networkTopologyFilename == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"errors": []map[string]interface{}{
				{
					"status": http.StatusNotFound,
					"title":  "no network topology data found",
					"detail": "no network topology data found for device " + device,
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
