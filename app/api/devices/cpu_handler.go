package devices

import (
	"github.com/ahmadateya/flotta-webapp-backend/helpers"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/analysis"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/logparser"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func (h *Handler) GetCPUTempData(c *gin.Context) {
	device := c.Param("device")
	// number of lines to read from log file
	lines := c.Query("lines")
	if lines == "" {
		lines = "50"
	}

	// read number of line n from the end of log file
	path := filepath.Join(logparser.LogDir, device, "cputemp")
	raw, err := logparser.ReadLogFileRaw(path, lines)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error reading cpu temp data",
			err.Error(),
		)
		return
	}
	// parse raw log file into proper structs
	parsedData, err := logparser.ParseCPUTempRawLines(raw)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error parsing cpu temp data",
			err.Error(),
		)
		return
	}

	// analyze the data to get insights
	cpuData, err := analysis.GetCPUAvgTempOverTheDay(parsedData)
	if err != nil {
		helpers.FormatErrorMessage(c,
			http.StatusInternalServerError,
			"error analyze temp data",
			err.Error(),
		)
		return
	}

	// return response
	c.JSON(http.StatusOK, gin.H{
		"data": map[string]interface{}{
			"labels":  analysis.HoursInDay,
			"degrees": cpuData,
		},
	})
}
