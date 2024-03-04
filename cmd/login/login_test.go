/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package login

import (
	"testing"
)

func TestNewLoginCmd(t *testing.T) {
	cmd := NewLoginCmd()
	if cmd == nil {
		t.Error("Expected non-nil command")
	}
}

func TestLoginCmd_Run(t *testing.T) {
	l := &loginCmd{}
	err := l.run(nil, nil)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
