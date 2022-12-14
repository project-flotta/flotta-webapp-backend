package analysis

import (
	"fmt"
	"github.com/ahmadateya/flotta-webapp-backend/pkg/logparser"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestGetCPUAvgTempOverTheDay(t *testing.T) {
	type args struct {
		raw []logparser.CPUTempParsedLine
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "TestParseCPUTempRawLines",
			args: args{
				raw: []logparser.CPUTempParsedLine{
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
				},
			},
			want:    []string{"40.00", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0", "0"},
			wantErr: false,
		},
		{
			name: "TestParseCPUTempRawLinesWithEmptyData",
			args: args{
				raw: []logparser.CPUTempParsedLine{},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCPUAvgTempOverTheDay(tt.args.raw)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, got, 24, fmt.Sprintf("Expected 24 avg degrees in day, got %d", len(got)))
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCPUAvgTempOverTheDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAvgTempDataInDay(t *testing.T) {
	tempData := map[string][]float64{
		"00": {50, 30},
	}
	avgTempDataInDay := getAvgTempDataInDay(tempData)
	assert.Equal(t, avgTempDataInDay["00"], "40.00", "Expected 40.00, got %s", avgTempDataInDay["00"])
}

func TestGetAvgTempInHour(t *testing.T) {
	temps := []float64{50, 50}
	avgTemp := getAvgTempInHour(temps)
	assert.Equal(t, avgTemp, "50.00", "Expected 50.00, got %s", avgTemp)
}

func TestGetTempMapData(t *testing.T) {
	tempData := getTempMapData()
	assert.Equal(t, len(tempData), 24, "Expected 24 hours, got %d", len(tempData))
}

func TestGetCPUAvgTempOverTheDayWithEmptyData(t *testing.T) {
	var logLines []logparser.CPUTempParsedLine
	avgTemp, err := GetCPUAvgTempOverTheDay(logLines)
	assert.Equal(t, err, fmt.Errorf("cant get avg temp over the day, log lines is empty"), "Expected empty data error, got %s", err)
	assert.Nil(t, avgTemp, "Expected nil, got %s", avgTemp)
}
