package analysis

import (
	"github.com/ahmadateya/flotta-webapp-backend/pkg/logparser"
	"testing"
)

func TestGetCPUAvgTempOverTheDay(t *testing.T) {
	logLines := []logparser.CPUTempParsedLine{
		{
			LogDate: "2020-01-01",
			LogTime: "00:00:00",
			Data: logparser.CPUTempData{
				Temp: 50,
			},
		},
		{
			LogDate: "2020-01-01",
			LogTime: "00:00:00",
			Data: logparser.CPUTempData{
				Temp: 30,
			},
		},
	}
	avgTemp, _ := GetCPUAvgTempOverTheDay(logLines)
	if len(avgTemp) != 24 {
		t.Errorf("Expected 24 avg degrees in day, got %d", len(avgTemp))
	}
	if avgTemp[0] != "40.00" {
		t.Errorf("Expected 40.00 in 00 hour, got %s", avgTemp[0])
	}
}

func TestGetAvgTempDataInDay(t *testing.T) {
	tempData := map[string][]float64{
		"00": {50, 30},
	}
	avgTempDataInDay := getAvgTempDataInDay(tempData)
	if avgTempDataInDay["00"] != "40.00" {
		t.Errorf("Expected 40.00, got %s", avgTempDataInDay["00"])
	}
}

func TestGetAvgTempInHour(t *testing.T) {
	temps := []float64{50, 50}
	avgTemp := getAvgTempInHour(temps)
	if avgTemp != "50.00" {
		t.Errorf("Expected 50.00, got %s", avgTemp)
	}
}

func TestGetTempMapData(t *testing.T) {
	tempData := getTempMapData()
	if len(tempData) != 24 {
		t.Errorf("Expected 24 hours, got %d", len(tempData))
	}
}

func TestGetCPUAvgTempOverTheDayWithEmptyData(t *testing.T) {
	var logLines []logparser.CPUTempParsedLine
	_, err := GetCPUAvgTempOverTheDay(logLines)
	if err == nil {
		t.Errorf("Expected error, got %v", err)
	}
}
