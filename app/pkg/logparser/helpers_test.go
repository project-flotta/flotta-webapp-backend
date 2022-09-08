package logparser

import (
	"testing"
)

func TestReadLogFileRaw(t *testing.T) {
	type args struct {
		dirPath string
		lines   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestReadLogFileRaw no such directory",
			args: args{
				dirPath: "test",
				lines:   "10",
			},
			wantErr: true,
		},
		{
			name: "TestReadLogFileRaw",
			args: args{
				dirPath: "testdata/device-name/cputemp",
				lines:   "10",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ReadLogFileRaw(tt.args.dirPath, tt.args.lines)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadLogFileRaw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLatestModifiedFile(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestGetLatestModifiedFile no such file or directory",
			args: args{
				dir: "test",
			},
			want:    "test",
			wantErr: true,
		},
		{
			name: "TestGetLatestModifiedFile file exists",
			args: args{
				dir: "testdata/device-name/cputemp",
			},
			want:    "cputemp.log",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLatestModifiedFile(tt.args.dir)
			if tt.wantErr {
				if (err != nil) != tt.wantErr {
					t.Errorf("getLatestModifiedFile() error = %v, wantErr %v", err, tt.wantErr)
				}
			} else {
				if got != tt.want {
					t.Errorf("getLatestModifiedFile() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestConvertAddressToString(t *testing.T) {
	type args struct {
		i []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestConvertAddressToString",
			args: args{
				i: []interface{}{"192", "168", "1", "1"},
			},
			want: "192.168.1.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertAddressToString(tt.args.i); got != tt.want {
				t.Errorf("convertAddressToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
