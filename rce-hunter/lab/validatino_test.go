package lab

import "testing"

func TestValidateUserName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid Username",
			input:    "mahdi123",
			expected: true,
		},
		{
			name:     "Empty Username",
			input:    "",
			expected: false,
		},
		{
			name:     "Too short",
			input:    "ab",
			expected: false,
		},
		{
			name:     "Too long",
			input:    "thisusernameiswaytoolong",
			expected: false,
		},
		{
			name:     "Whitespace Only",
			input:    "	",
			expected: false,
		},
		{
			name:     "Invalid Format",
			input:    "m@hdi123",
			expected: false,
		},
		{
			name:     "Underscore",
			input:    "mahdi_123",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateUsername(tt.input)

			if result.Valid != tt.expected {
				t.Errorf(
					"Expected Valid = %v, Got Valid = %v",
					tt.expected,
					result.Valid,
				)
			}
		})
	}
}

func TestValidationErrors(t *testing.T) {
	result := ValidateUsername("ab")

	if result.Valid {
		t.Error("Expected invalid Username")
	}
	if len(result.Errors) == 0 {
		t.Error("Expected at Least One Validation Error")
	}
}
