/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package login

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestNewLoginCmd(t *testing.T) {
	tests := []struct {
		name string
		want *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoginCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoginCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
