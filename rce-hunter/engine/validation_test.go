package engine

import (
	"strings"
	"testing"
)

func TestValidationSimulation(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Valid Username",
			input:    "mahdi123",
			expected: "PASS",
		},
		{
			name:     "Invalid Username",
			input:    "ab\n",
			expected: "REJECTED",
		},
		{
			name:     "Invalid Format",
			input:    "m@hdi123\n",
			expected: "REJECTED",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			simulation := &ValidationSimulation{
				Lab: Lab{
					ID:   "val-001",
					Name: "Validation Test Lab",
				},
				Input: strings.NewReader(tt.input),
			}
			result, err := simulation.Run()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if result.Status != tt.expected {
				t.Errorf("Expected Status = %v, Got Status = %v", tt.expected, result.Status)
			}
		})
	}
}
