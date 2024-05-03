package main

import "testing"

func Test_timeFunc(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := timeFunc(); got != tt.want {
				t.Errorf("timeFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
