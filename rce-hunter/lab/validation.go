package lab

import (
	"regexp"
	"strings"
)

type ValidationResult struct {
	Input  string
	Valid  bool
	Errors []string
}

func ValidateUsername(input string) ValidationResult {
	result := ValidationResult{
		Input:  input,
		Valid:  true,
		Errors: []string{},
	}

	if strings.TrimSpace(input) == "" {
		result.Valid = false
		result.Errors = append(result.Errors, "Username Is Required")
	}

	if len(input) < 3 {
		result.Valid = false
		result.Errors = append(result.Errors, "Username Must Be At Least 3 Characters")
	}

	if len(input) > 20 {
		result.Valid = false
		result.Errors = append(result.Errors, "Username Must Not Be Exceed 20 Characters")
	}

	usernamePattern := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

	if !usernamePattern.MatchString(input) && strings.TrimSpace(input) != "" {
		result.Valid = false
		result.Errors = append(result.Errors, "Username Contains Invalid Characters")
	}
	return result
}
