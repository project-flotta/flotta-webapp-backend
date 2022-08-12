package devices

import (
	"encoding/json"
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/helpers"
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

func (h *Handler) GetNetworkData(c *gin.Context) {
	device := c.Param("device")

	// get network topology data from S3
	client := s3.InitS3Client()
	networkTopologyFilename, err := h.S3.GetMostRecentObjectNameInFolder(device + "/network")
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusNotFound,
			"error getting network data",
			err.Error(),
		)
		return
	}

	// download network topology file from S3
	filename := networkTopologyFilename[strings.LastIndex(networkTopologyFilename, "/")+1:]
	err = h.S3.DownloadObject(filename, networkTopologyFilename)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusNotFound,
			"error getting network data",
			err.Error(),
		)
		return
	}

	// read contents of network topology file from S3
	objContent, err := client.ReadObject(networkTopologyFilename)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusNotFound,
			"error getting network data",
			err.Error(),
		)
		return
	}

	fmt.Printf("=================== objContent: %s\n", objContent)
	// parse objContent data to JSON
	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(objContent), &jsonMap)

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": jsonMap,
	})
}
