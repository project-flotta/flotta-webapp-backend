package devices

import (
	"github.com/ahmadateya/flotta-webapp-backend/helpers"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/logparser"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetCPUTempData(c *gin.Context) {
	device := c.Param("device")
	// number of lines to read from log file
	lines := c.Query("lines")
	if lines == "" {
		lines = "5"
	}

	// read number of line n from the end of log file
	raw, err := logparser.ReadLogFileRaw(device+"/cputemp", lines)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error reading cpu temp data",
			err.Error(),
		)
		return
	}
	// parse raw log file into proper structs
	cpuData, err := logparser.ParseCPUTempRawLines(raw)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error parsing cpu temp data",
			err.Error(),
		)
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": cpuData,
	})
}
