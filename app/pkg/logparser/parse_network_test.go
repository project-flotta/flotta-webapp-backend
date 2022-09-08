package logparser

import (
	"reflect"
	"testing"
)

func TestParseNetworkRawLines(t *testing.T) {
	raw := "2022/08/04 22:05:26 {\"DestinationAddress\":[142,250,201,46],\"Hops\":[{\"Success\":true,\"Address\":[172,18,0,1],\"Host\":\"ateya-home.local.\",\"N\":52,\"ElapsedTime\":254297,\"TTL\":1}]}\n"
	got, err := ParseNetworkRawLines(raw)
	uuid := got[0].Data.Hops[0].Uuid
	type args struct {
		raw string
	}
	tests := []struct {
		name    string
		args    args
		want    []NetworkParsedLine
		wantErr bool
	}{
		{
			name: "TestParseNetworkRawLines",
			want: []NetworkParsedLine{
				{
					LogDate: "2022/08/04",
					LogTime: "22:05:26",
					Data: NetworkData{
						DestinationAddress: "142.250.201.46",
						Hops: []Hop{
							{
								Success:     true,
								Address:     "172.18.0.1",
								Host:        "ateya-home.local.",
								N:           52,
								ElapsedTime: 254297,
								TTL:         1,
								Uuid:        uuid,
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNetworkRawLines() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetworkRawLines() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseNetworkSingleLine(t *testing.T) {
	lineData := []string{
		"2022/08/04",
		"22:05:26",
		"{\"DestinationAddress\":[142,250,201,46],\"Hops\":[{\"Success\":true,\"Address\":[172,18,0,1],\"Host\":\"ateya-home.local.\",\"N\":52,\"ElapsedTime\":254297,\"TTL\":1}]}",
	}
	got, err := parseNetworkSingleLine(lineData)
	uuid := got.Data.Hops[0].Uuid
	tests := []struct {
		name    string
		want    NetworkParsedLine
		wantErr bool
	}{
		{
			name: "TestParseNetworkSingleLine",
			want: NetworkParsedLine{
				LogDate: "2022/08/04",
				LogTime: "22:05:26",
				Data: NetworkData{
					DestinationAddress: "142.250.201.46",
					Hops: []Hop{
						{
							Success:     true,
							Address:     "172.18.0.1",
							Host:        "ateya-home.local.",
							N:           52,
							ElapsedTime: 254297,
							TTL:         1,
							Uuid:        uuid,
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseNetworkSingleLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseNetworkSingleLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}
