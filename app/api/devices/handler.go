package devices

import (
	"net/http"
	"strings"
	"github.com/ahmadateya/flotta-webapp-backend/helpers"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/s3"
	"github.com/gin-gonic/gin"
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

func (h *Handler) GetDeviceData(c *gin.Context) {
	device := c.Param("device")

	// download network log file from S3
	err := h.downloadLogFile(device + "/network")
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error getting network data",
			err.Error(),
		)
		return
	}

	// download cputemp log file from S3
	err = h.downloadLogFile(device + "/cputemp")
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error getting CPU temperature data",
			err.Error(),
		)
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": "success",
	})
}

func (h *Handler) downloadLogFile(device string) error {
	// get latest object in folder from S3
	filename, err := h.S3.GetMostRecentObjectNameInFolder(device)
	if err != nil {
		return err
	}

	// download file from S3
	err = h.S3.DownloadObject(filename)
	if err != nil {
		return err
	}
	return nil
}
