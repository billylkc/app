package util

import (
	"reflect"
	"testing"
	"time"
)

func TestParseDateInput(t *testing.T) {
	type args struct {
		s    string
		freq string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// {
		// 	name: "Next week",
		// 	args: args{
		// 		s:    "2021-04-03",
		// 		freq: "w",
		// 	},
		// 	want: "2021-04-05",
		// },
		{
			name: "Back 5 days",
			args: args{
				s:    "5",
				freq: "d",
			},
			want: time.Now().AddDate(0, 0, -5).Format("2006-01-02"),
		},
		{
			name: "today",
			args: args{
				s:    time.Now().Format("2006-01-02"),
				freq: "d",
			},
			want: time.Now().Format("2006-01-02"),
		},
		{
			name: "Normal day in YYYY-MM-DD",
			args: args{
				s:    "2021-03-01",
				freq: "d",
			},
			want: "2021-03-01",
		},
		{
			name: "Current month",
			args: args{
				s:    "2021-03-11",
				freq: "m",
			},
			want: "2021-03-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDateInput(tt.args.s, tt.args.freq)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseDateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseDateInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGenerateDate(t *testing.T) {
	type args struct {
		start string
		end   string
		freq  string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Simple test - day",
			args: args{
				start: "2019-02-20",
				end:   "2019-02-22",
				freq:  "d",
			},
			want: []string{
				"2019-02-20",
				"2019-02-21",
			},
		},
		{
			name: "Simple test - week",
			args: args{
				start: "2021-05-03",
				end:   "2021-05-17",
				freq:  "w",
			},
			want: []string{
				"2021-05-03",
				"2021-05-10",
			},
		},
		{
			name: "Simple test - month",
			args: args{
				start: "2021-02-01",
				end:   "2021-05-31",
				freq:  "m",
			},
			want: []string{
				"2021-02-01",
				"2021-03-01",
				"2021-04-01",
				"2021-05-01",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateDate(tt.args.start, tt.args.end, tt.args.freq)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateDate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
