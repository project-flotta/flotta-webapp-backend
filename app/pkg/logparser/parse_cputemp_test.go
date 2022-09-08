package logparser

import "testing"

func TestParseCPUTempRawLines(t *testing.T) {
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    []CPUTempParsedLine
		wantErr bool
	}{
		{
			name: "TestParseCPUTempRawLines",
			args: args{
				raw: "2020-01-01 00:00:00 {\"Temp\":50.0}\n2020-01-01 00:00:01 {\"Temp\":40.0}\n2020-01-01 00:00:02 {\"Temp\":30.0}\n",
			},
			want: []CPUTempParsedLine{
				{
					LogDate: "2020-01-01",
					LogTime: "00:00:00",
					Data: CPUTempData{
						Temp: 50.0,
					},
				},
				{
					LogDate: "2020-01-01",
					LogTime: "00:00:01",
					Data: CPUTempData{
						Temp: 40.0,
					},
				},
				{
					LogDate: "2020-01-01",
					LogTime: "00:00:02",
					Data: CPUTempData{
						Temp: 30.0,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCPUTempRawLines(tt.args.raw)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCPUTempRawLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("ParseCPUTempRawLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseCPUTempSingleLine(t *testing.T) {
	type args struct {
		lineData []string
	}
	tests := []struct {
		name    string
		args    args
		want    CPUTempParsedLine
		wantErr bool
	}{
		{
			name: "TestParseCPUTempSingleLine",
			args: args{
				lineData: []string{"2020-01-01", "00:00:00", "{\"Temp\": 50.0}"},
			},
			want: CPUTempParsedLine{
				LogDate: "2020-01-01",
				LogTime: "00:00:00",
				Data: CPUTempData{
					Temp: 50.0,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCPUTempSingleLine(tt.args.lineData)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCPUTempSingleLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.LogDate != tt.want.LogDate {
				t.Errorf("parseCPUTempSingleLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
