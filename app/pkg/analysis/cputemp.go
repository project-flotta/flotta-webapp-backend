package analysis

import (
	"fmt"
	logparser "github.com/ahmadateya/flotta-webapp-backend/pkg/logparser"
)

// HoursInDay slice of hours in day
var HoursInDay = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}

func GetCPUAvgTempOverTheDay(logLines []logparser.CPUTempParsedLine) ([]string, error) {
	if len(logLines) == 0 {
		return nil, fmt.Errorf("cant get avg temp over the day, log lines is empty")
	}
	// get the first date in the data
	firstDate := logLines[0].LogDate
	// collect all temp data in the same day in a map with key as the hour of the day
	tempData := getTempMapData()
	for _, d := range logLines {
		// check if the date is the same as the first date
		if d.LogDate == firstDate {
			// get the hour of the day
			hour := d.LogTime[:2]
			// append the temp data to the hour key
			tempData[hour] = append(tempData[hour], d.Data.Temp)
		}
	}
	// calculate the average temp for each hour
	avgTempDataInDay := getAvgTempDataInDay(tempData)
	sorted := sortAvgTempDataInDay(avgTempDataInDay)
	return sorted, nil
}

func getTempMapData() map[string][]float64 {
	return map[string][]float64{
		"00": {}, "01": {}, "02": {}, "03": {},
		"04": {}, "05": {}, "06": {}, "07": {},
		"08": {}, "09": {}, "10": {}, "11": {},
		"12": {}, "13": {}, "14": {}, "15": {},
		"16": {}, "17": {}, "18": {}, "19": {},
		"20": {}, "21": {}, "22": {}, "23": {},
	}
}

func getAvgTempDataInDay(tempDataInDay map[string][]float64) map[string]string {
	avgTempDataInDay := make(map[string]string, 24)
	// calculate the average temp for each hour
	for hour, temps := range tempDataInDay {
		// calculate the average temp
		avgTempDataInDay[hour] = getAvgTempInHour(temps)
	}
	return avgTempDataInDay
}

func getAvgTempInHour(temps []float64) string {
	if len(temps) > 0 {
		// calculate the average temp
		var avg float64
		for _, t := range temps {
			avg += t
		}
		avg /= float64(len(temps))
		return fmt.Sprintf("%.2f", avg)
	} else {
		return "0"
	}
}

// sort avgTempDataInDay by hour
func sortAvgTempDataInDay(m map[string]string) []string {
	var sorted []string
	for _, h := range HoursInDay {
		sorted = append(sorted, m[h])
	}
	return sorted
}
