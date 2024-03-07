/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package login

import (
	"testing"
)

func TestNewLoginCmd(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test New Login Cmd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lgn := &loginCmd{}
			err := lgn.run(nil, nil)
			if err != nil {
				t.Errorf("NewLoginCmd() error = %v", err)
			}
		})
	}
}
