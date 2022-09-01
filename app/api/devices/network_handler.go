package devices

import (
	"net/http"

	"github.com/ahmadateya/flotta-webapp-backend/helpers"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/logparser"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetNetworkData(c *gin.Context) {
	device := c.Param("device")
	// number of lines to read from log file
	numOfLines := c.Query("logs")
	if numOfLines == "" || numOfLines == "null" {
		numOfLines = "2"
	}

	// read number of line n from the end of log file
	raw, err := logparser.ReadLogFileRaw(device+"/network", numOfLines)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error reading network data",
			err.Error(),
		)
		return
	}
	// parse raw log file into proper structs
	netData, err := logparser.ParseNetworkRawLines(raw)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error parsing network data",
			err.Error(),
		)
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": netData,
	})
}
