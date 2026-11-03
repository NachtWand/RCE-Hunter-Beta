package report

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rce-hunter/checks"
	"rce-hunter/engine"
)

func PrintJson(result engine.Result) error {
	data, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return fmt.Errorf("encode report: %w", err)
	}
	fmt.Println(string(data))
	return nil
}

func SaveJson(result engine.Result) (string, error) {
	if result.LabID == "" || strings.ContainsAny(result.LabID, `/\`) {
		return "", fmt.Errorf("Invalid Lab ID")
	}
	if err := os.MkdirAll("report", 0755); err != nil {
		return "", fmt.Errorf("create report directory: %w", err)
	}

	data, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return "", fmt.Errorf("encode report: %w", err)
	}

	path := filepath.Join("report", result.LabID+".json")

	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", fmt.Errorf("save report: %w", err)
	}

	return path, nil
}

func AppendHistory(result engine.Result) error {
	const historyPath = "report/history.json"

	if err := os.MkdirAll("report", 0755); err != nil {
		return fmt.Errorf("create report directory: %w", err)
	}

	var history []engine.Result

	data, err := os.ReadFile(historyPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Read History: %w", err)
	}

	if err == nil {
		if err := json.Unmarshal(data, &history); err != nil {
			return fmt.Errorf("decode history: %w", err)
		}
	}

	history = append(history, result)

	data, err = json.MarshalIndent(history, "", "  ")
	if err != nil {
		return fmt.Errorf("endoce history: %w", err)
	}

	if err := os.WriteFile(historyPath, data, 0644); err != nil {
		return fmt.Errorf("write history: %w", err)
	}

	return nil

}

func ShowHistory() error {
	const historyPath = "report/history.json"

	data, err := os.ReadFile(historyPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No Test History Found")
			return nil
		}
		return fmt.Errorf("read history: %w", err)
	}

	var history []engine.Result

	if err := json.Unmarshal(data, &history); err != nil {
		return fmt.Errorf("decode history: %w", err)
	}

	if len(history) == 0 {
		fmt.Println("No Test History Found")
		return nil
	}

	fmt.Println("\n=== Test History ===")

	for i, result := range history {
		fmt.Printf("\n[%d] %s\n	Lab ID: %s\n	Status: %s\n	Message: %s\n	Executed_At: %s\n",
			i+1,
			result.LabName,
			result.LabID,
			result.Status,
			result.Message,
			result.ExecutedAt.Format("2006-01-02 15:04:05 UTC"),
		)
	}

	fmt.Printf("\nTotal Runs: %d\n", len(history))
	return nil
}

func SaveAnalysis(result checks.InjectionResult) (string, error) {
	if err := os.MkdirAll("report", 0755); err != nil {
		return "", fmt.Errorf("Create Report Directory: %w", err)
	}

	data, err := json.MarshalIndent(result, "", " ")
	if err != nil {
		return "", fmt.Errorf("Encode Analysis Report: %w", err)
	}

	path := filepath.Join("report", "analysis--01.json")
	if err := os.WriteFile(path, data, 0644); err != nil {
		return "", fmt.Errorf("Save Analysis Report: %w", err)
	}

	return path, nil
}
