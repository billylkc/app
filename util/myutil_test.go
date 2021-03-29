package util

import (
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
