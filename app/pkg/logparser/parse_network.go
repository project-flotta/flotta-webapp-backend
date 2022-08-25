package logparser

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
	"time"
)

type NetworkParsedLine struct {
	LogDate string      `json:"log_date"`
	LogTime string      `json:"log_time"`
	Data    NetworkData `json:"data"`
}

type NetworkData struct {
	DestinationAddress string `json:"destination_address"`
	Hops               []Hop  `json:"hops"`
}

type Hop struct {
	Success     bool          `json:"success"`
	Uuid        string        `json:"uuid"`
	Address     string        `json:"address"`
	Host        string        `json:"host"`
	N           float64       `json:"n"`
	ElapsedTime time.Duration `json:"elapsed_time"`
	TTL         float64       `json:"ttl"`
}

func ParseNetworkRawLines(raw string) ([]NetworkParsedLine, error) {
	// split raw string into lines
	linesRegex := regexp.MustCompile("\n")
	lines := linesRegex.Split(raw, -1)

	var parsedLines []NetworkParsedLine
	for _, l := range lines {
		// split raw string into date, time and data
		if len(l) > 0 {
			lineData := strings.Split(l, " ")
			parsedLine, err := parseNetworkSingleLine(lineData)
			if err != nil {
				return nil, err
			}
			parsedLines = append(parsedLines, parsedLine)
		}
	}
	return parsedLines, nil
}

func parseNetworkSingleLine(lineData []string) (NetworkParsedLine, error) {
	var parsedLine NetworkParsedLine
	var netData map[string]interface{}
	var err error
	parsedLine.LogDate = lineData[0]
	parsedLine.LogTime = lineData[1]

	// Unmarshal the JSON into a struct
	err = json.Unmarshal([]byte(lineData[2]), &netData)
	if err != nil {
		return parsedLine, fmt.Errorf("error unmarshaling network data: %v", err)
	}

	parsedLine.Data.fillIntoStruct(netData)
	return parsedLine, nil
}

func (s *NetworkData) fillIntoStruct(netData map[string]interface{}) {
	s.DestinationAddress = convertAddressToString(netData["DestinationAddress"].([]interface{}))
	iHop := netData["Hops"].([]interface{})
	for _, h := range iHop {
		s.Hops = append(s.Hops, fillHopDataIntoStruct(h.(map[string]interface{})))
	}
}

func fillHopDataIntoStruct(hopData map[string]interface{}) Hop {
	hop := Hop{}
	hop.Host = fmt.Sprintf("%v", hopData["Host"])
	hop.Address = convertAddressToString(hopData["Address"].([]interface{}))
	hop.Success = hopData["Success"].(bool)
	hop.N = hopData["N"].(float64)
	hop.ElapsedTime = time.Duration(hopData["ElapsedTime"].(float64))
	hop.TTL = hopData["TTL"].(float64)
	hop.Uuid = uuid.New().String()
	return hop
}
