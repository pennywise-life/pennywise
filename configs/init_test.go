package configs

import "testing"

func TestInitConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Test Init Config",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfig()
		})
	}
}
