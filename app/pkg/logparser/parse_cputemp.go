package logparser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

type CPUTempData struct {
	Temp float64 `json:"temp"`
}

type CPUTempParsedLine struct {
	LogDate string      `json:"log_date"`
	LogTime string      `json:"log_time"`
	Data    CPUTempData `json:"data"`
}

func ParseCPUTempRawLines(raw string) ([]CPUTempParsedLine, error) {
	// split raw string into lines
	linesRegex := regexp.MustCompile("\n")
	lines := linesRegex.Split(raw, -1)

	var parsedLines []CPUTempParsedLine
	for _, l := range lines {
		// split raw string into date, time and data
		if len(l) > 0 {
			lineData := strings.Split(l, " ")
			parsedLine, err := parseCPUTempSingleLine(lineData)
			if err != nil {
				return nil, err
			}
			parsedLines = append(parsedLines, parsedLine)
		}
	}
	return parsedLines, nil
}

func parseCPUTempSingleLine(lineData []string) (CPUTempParsedLine, error) {
	var parsedLine CPUTempParsedLine
	var cpuTempData map[string]interface{}
	var err error
	parsedLine.LogDate = lineData[0]
	parsedLine.LogTime = lineData[1]

	// Unmarshal the JSON into a struct
	err = json.Unmarshal([]byte(lineData[2]), &cpuTempData)
	if err != nil {
		return parsedLine, fmt.Errorf("error unmarshaling cpu temp data: %v", err)
	}
	parsedLine.Data.Temp = cpuTempData["Temp"].(float64)
	return parsedLine, nil
}
