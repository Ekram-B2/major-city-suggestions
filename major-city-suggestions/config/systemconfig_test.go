package config

import "testing"

func Test_getConfigPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getConfigPath(); got != tt.want {
				t.Errorf("getConfigPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
