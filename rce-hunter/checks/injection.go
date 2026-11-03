package checks

import "strings"

type Finding struct {
	Rule        string `json:"rule"`
	Description string `json:"description"`
}

type InjectionResult struct {
	Input          string    `json:"input"`
	Classfication  string    `json:"classfication"`
	EvidenceLevel  string    `json:"evidence_level"`
	Findings       []Finding `json:"findings"`
	Limitations    []string  `json:"limitations"`
	Recommendation string    `json:"recommendation"`
	SafeHandling   string    `json:"safe_handling"`
}

func AnalyzeInput(input string) InjectionResult {
	result := InjectionResult{
		Input:         input,
		Classfication: "NORMAL",
		EvidenceLevel: "NONE",
		Findings:      []Finding{},
		Limitations: []string{
			"Character matching alone cannot confirm command injection.",
			"Application behavior and input handling must be reviewed separately.",
		},
		Recommendation: "Review how the application validates and processes this input.",
		SafeHandling:   "Treat input as data and never execute it as an operating system command.",
	}

	if strings.TrimSpace(input) == "" {
		result.Findings = append(result.Findings, Finding{
			Rule:        "EMPTY_INPUT",
			Description: "Input is empty or contains only whitespace",
		})
	}

	for _, char := range input {
		if char == '\n' || char == '\t' || char == '\r' {
			result.Findings = append(result.Findings, Finding{
				Rule:        "CONTROL_CHARACTER",
				Description: "Input Contains a Control Character",
			})
			break
		}
	}

	if strings.ContainsAny(input, ";&|`$") {
		result.Classfication = "SUSPICIOUS"
		result.EvidenceLevel = "LOW"

		result.Findings = append(result.Findings, Finding{
			Rule:        "SPECIAL_CHARACTER",
			Description: "Input contains characters that may require contextual security review.\n",
		})
	}

	if len(result.Findings) > 0 {
		result.Classfication = "SUSPICIOUS"
		result.EvidenceLevel = "LOW"

	}

	return result
}
