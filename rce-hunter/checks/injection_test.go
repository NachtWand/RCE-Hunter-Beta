package checks

import (
	"testing"
)

func TestAnalyzeInput(t *testing.T) {
	tests := []struct {
		name                 string
		input                string
		expected             string
		expectedFindingCount int
	}{
		{
			name:                 "Normal input",
			input:                "hello world",
			expected:             "NORMAL",
			expectedFindingCount: 0,
		},
		{
			name:                 "Suspicious input",
			input:                "hello; world",
			expected:             "SUSPICIOUS",
			expectedFindingCount: 1,
		},
		{
			name:                 "Empty input",
			input:                "",
			expected:             "SUSPICIOUS",
			expectedFindingCount: 1,
		},
		{
			name:                 "Input with spaces",
			input:                "hello world test",
			expected:             "NORMAL",
			expectedFindingCount: 0,
		},
		{
			name:                 "WhiteSpace Input",
			input:                " ",
			expected:             "SUSPICIOUS",
			expectedFindingCount: 1,
		},
		{
			name:                 "Tab Character",
			input:                "hello\tworld",
			expected:             "SUSPICIOUS",
			expectedFindingCount: 1,
		},
		{
			name:                 "New Line Character",
			input:                "hello\nworld",
			expected:             "SUSPICIOUS",
			expectedFindingCount: 1,
		},
		{
			name:                 "Multiple Rules",
			input:                "hello;\nworld",
			expected:             "SUSPICIOUS",
			expectedFindingCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AnalyzeInput(tt.input)

			if result.Classfication != tt.expected {
				t.Errorf(
					"expected classification %q, got %q",
					tt.expected,
					result.Classfication,
				)
			}

			if len(result.Findings) != tt.expectedFindingCount {
				t.Errorf(
					"expected %d findings, got %d",
					tt.expectedFindingCount,
					len(result.Findings),
				)
			}
		})
	}
}
