package main

import (
	"fmt"
	"os"
	"strings"

	"rce-hunter/checks"
	"rce-hunter/engine"
	"rce-hunter/lessons"
	"rce-hunter/report"
)

func main() {
	if len(os.Args) < 2 {
		showHelp()
		return
	}

	switch os.Args[1] {
	case "lessons":
		lessons.List()

	case "lesson":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . lesson <lesson-id>")
			return
		}
		if !lessons.Show(os.Args[2]) {
			fmt.Println("Lesson not found")
		}

	case "labs":
		listlabs()

	case "run":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . run <lab-id>")
			return
		}
		runLab(os.Args[2])

	case "version":
		fmt.Println("RCE Hunter version 1.0.0")

	case "history":
		if err := report.ShowHistory(); err != nil {
			fmt.Println("Could not Show History:", err)
		}
	case "analyze":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run . analyze <input>")
			return
		}

		input := strings.Join(os.Args[2:], " ")
		result := checks.AnalyzeInput(input)

		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("          RCE-HUNTER ANALYSIS")
		fmt.Println("========================================")
		fmt.Printf("Input          : %s\n", result.Input)
		fmt.Printf("Classification : %s\n", result.Classfication)
		fmt.Printf("Evidence Level : %s\n", result.EvidenceLevel)
		fmt.Println("----------------------------------------")

		fmt.Println("FINDINGS")
		if len(result.Findings) == 0 {
			fmt.Println("No Suspicious Patterns Detected")
		} else {
			for i, finding := range result.Findings {
				fmt.Printf(" [%d] Rule : %s\n", i+1, finding.Rule)
				fmt.Printf("Description : %s\n", finding.Description)
			}
		}

		fmt.Println("----------------------------------------")
		fmt.Println("LIMITATIONS")
		for i, limitation := range result.Limitations {
			fmt.Printf(" %d.%s\n", i+1, limitation)
		}

		fmt.Println("----------------------------------------")
		fmt.Printf("Recommendation : %s\n", result.Recommendation)
		fmt.Printf("Safe Handling  : %s\n", result.SafeHandling)
		fmt.Println("========================================")

		path, err := report.SaveAnalysis(result)
		if err != nil {
			fmt.Printf("\n[!].Failed To Save Analysis Report: %v\n", err)
			return
		}
		fmt.Printf("\n[+].Full Analysis Report Saved To: %s\n", path)

	default:
		showHelp()
	}
}

func listlabs() {
	fmt.Println("\n=== RCE Lab | Available Labs ===")

	for _, Lab := range engine.AvailableLabs() {
		fmt.Printf(
			"\nID: %s\nName: %s\nDescription: %s\n",
			Lab.ID,
			Lab.Name,
			Lab.Description,
		)
	}
}

func runLab(id string) {
	result, err := engine.Execute(id)

	if result.LabID != "" {
		if result.LabID == "val-001" {
			printValidationResult(result)
		} else {
			printResult(result)
		}

		path, saveErr := report.SaveJson(result)

		if saveErr != nil {
			fmt.Println("Could not save report:", saveErr)
		} else {
			fmt.Println("Report saved to:", path)
		}

		if saveErr := report.AppendHistory(result); saveErr != nil {
			fmt.Println("Could Not Update History:", saveErr)
		} else {
			fmt.Println("History Updated")
		}
	}

	if err != nil {
		fmt.Println("Lab execution failed:", err)
		return
	}
}

func printResult(result engine.Result) {
	if err := report.PrintJson(result); err != nil {
		fmt.Println("Could Not Print Report:", err)
	}
}

func showHelp() {
	fmt.Println("RCE Lab - Security Learning Platform")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  go run . lessons")
	fmt.Println("  go run . lesson <lesson-id>")
	fmt.Println("  go run . labs")
	fmt.Println("  go run . run <lab-id>")
	fmt.Println("  go run . version")
	fmt.Println("  go run . history")
	fmt.Println("Usage: go run . analyze <input>")
}

func printValidationResult(result engine.Result) {
	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("          RCE-HUNTER VALIDATION")
	fmt.Println("========================================")
	fmt.Printf("Input          : %s\n", result.Input)
	fmt.Printf("Validation     : %s\n", result.Status)

	if len(result.ValidationErrors) > 0 {
		fmt.Println("Validation Errors:")
		for i, err := range result.ValidationErrors {
			fmt.Printf(" [%d] %s\n", i+1, err)
		}
	} else {
		fmt.Println("\nUsername Accepted Successfully")
	}
	fmt.Println("========================================")
}
