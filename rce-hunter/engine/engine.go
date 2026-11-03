package engine

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"rce-hunter/checks"
	"rce-hunter/lab"
)

type Lab struct {
	ID          string
	Name        string
	Description string
}

type Result struct {
	LabID            string        `json:"lab_id"`
	LabName          string        `json:"lab_name"`
	Status           string        `json:"status"`
	Message          string        `json:"message"`
	Duration         time.Duration `json:"duration"`
	ExecutedAt       time.Time     `json:"executed_at"`
	Input            string        `json:"input,omitempty"`
	ValidationErrors []string      `json:"validation_errors,omitempty"`
}

type Runner interface {
	Run() (Result, error)
}

type HTTPSimulation struct {
	Lab Lab
}

type InjectionSimulation struct {
	Lab Lab
}

type ValidationSimulation struct {
	Lab   Lab
	Input io.Reader
}

func AvailableLabs() []Lab {
	return []Lab{
		{
			ID:          "http--01",
			Name:        "HTTP Health Check",
			Description: "A simple lab to check the health of an HTTP endpoint.",
		},
		{
			ID:          "rce--01",
			Name:        "RCE Behavior Simulation",
			Description: "A lab to simulate Remote Code Execution (RCE) behavior.",
		},
		{
			ID:          "inj-001",
			Name:        "Command Injection Simulation",
			Description: "Learn how suspicious input can be detected safely.",
		},
		{
			ID:          "val-001",
			Name:        "Input Validation Simulation",
			Description: "A lab to demonstrate input validation techniques.",
		},
	}
}

func NewRunner(id string) (Runner, error) {
	for _, lab := range AvailableLabs() {
		if lab.ID != id {
			continue
		}

		switch id {
		case "inj-001":
			return &InjectionSimulation{Lab: lab}, nil
		case "val-001":
			return &ValidationSimulation{Lab: lab}, nil
		default:
			return &HTTPSimulation{Lab: lab}, nil
		}
	}

	return nil, fmt.Errorf("Unknown lab: %s", id)
}

func (s *HTTPSimulation) Run() (Result, error) {
	startTime := time.Now()

	result := Result{
		LabID:      s.Lab.ID,
		LabName:    s.Lab.Name,
		Status:     "success",
		Message:    "Lab executed successfully",
		ExecutedAt: startTime,
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	var endpoint string
	var expected string

	switch s.Lab.ID {
	case "http--01":
		endpoint = "http://localhost:8080/health"
		expected = "LAB_OK"
	case "rce--01":
		endpoint = "http://localhost:8080/simulate?input=LAB_TEST"
		expected = "SIMULATED_UNSAFE_PATH: LAB_TEST"
	default:
		return Result{}, fmt.Errorf("lab with ID %s not implemented", s.Lab.ID)
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return result, err
	}

	resp, err := client.Do(req)
	if err != nil {
		result.Message = "Could Not Reach The Local Lab Server. Please ensure the server is running and accessible."
		result.Duration = time.Since(startTime)
		return result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		result.Message = fmt.Sprintf("Unexpected status code: %d", resp.StatusCode)
		result.Duration = time.Since(startTime)

		return result, fmt.Errorf("%s", result.Message)
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 4096))
	if err != nil {
		result.Message = "Failed to read lab response"
		result.Duration = time.Since(startTime)
		return result, err
	}

	actual := strings.TrimSpace(string(body))

	if actual != expected {
		result.Message = fmt.Sprintf(
			"Response mismatch: expected %q, received %q",
			expected,
			actual,
		)
		result.Duration = time.Since(startTime)

		return result, fmt.Errorf("%s", result.Message)
	}

	result.Status = "PASS"
	result.Message = "Expected Lab Behavior Confirmed"
	result.Duration = time.Since(startTime)

	return result, nil
}

func (s *InjectionSimulation) Run() (Result, error) {
	start := time.Now()

	result := Result{
		LabID:      s.Lab.ID,
		LabName:    s.Lab.Name,
		Status:     "PASS",
		ExecutedAt: start.UTC(),
	}

	analysis := checks.AnalyzeInput("sample;input")

	if analysis.Classfication != "SUSPICIOUS" {
		result.Status = "Failed"
		result.Message = "Expected Suspiciuos Input Was Not Detected"
		result.Duration = time.Since(start)
		return result, fmt.Errorf("%s", result.Message)
	}

	result.Message = "Suspicious input classified by the educational analysis engine"
	result.Duration = time.Since(start)

	return result, nil
}

func (s *ValidationSimulation) Run() (Result, error) {
	start := time.Now()

	result := Result{
		LabID:      s.Lab.ID,
		LabName:    s.Lab.Name,
		Status:     "PASS",
		ExecutedAt: start.UTC(),
	}

	inputreader := s.Input
	if inputreader == nil {
		inputreader = os.Stdin
	}

	reader := bufio.NewReader(inputreader)
	fmt.Print("Enter Your Username: ")

	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		result.Duration = time.Since(start)
		return result, err
	}
	input = strings.TrimRight(input, "\r\n")

	validationResult := lab.ValidateUsername(input)
	result.Input = validationResult.Input
	result.ValidationErrors = validationResult.Errors

	if !validationResult.Valid {
		result.Status = "REJECTED"
		result.Message = strings.Join(validationResult.Errors, "; ")
		result.Duration = time.Since(start)
		return result, nil
	}
	result.Message = "Username Accepted Succesfully:  " + validationResult.Input
	result.Duration = time.Since(start)

	return result, nil
}

func Execute(id string) (Result, error) {
	runner, err := NewRunner(id)
	if err != nil {
		return Result{}, err
	}
	return runner.Run()
}
