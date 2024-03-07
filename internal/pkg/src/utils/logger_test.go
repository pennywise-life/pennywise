package utils

import (
	"reflect"
	"testing"
)

func TestNewLog(t *testing.T) {
	tests := []struct {
		name string
		want *Log
	}{
		{
			name: "Test New Log",
			want: NewLog(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.want ;!reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
