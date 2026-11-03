package lessons

import "fmt"

type Lesson struct {
	ID          string
	Title       string
	Description string
	Level       string
	Content     string
	Objective   string
}

func GetAll() []Lesson {
	return []Lesson{
		{
			ID:          "rce--01",
			Title:       "Remote Code Execution (RCE) Fundementals",
			Description: "Learn the basics of Remote Code Execution (RCE) vulnerabilities, how they work, and how to identify them in web applications.",
			Level:       "Beginner",
			Objective:   "Understand how remote code execution vulnerabilities work.",
			Content: `
Remote Code Execution (RCE) is a security vulnerability
that may allow an attacker to execute code on a remote system.

Common root causes include:
- Unsafe command construction
- Improper input handling
- Vulnerable application components

Defensive measures:
- Avoid invoking operating system commands when possible
- Validate input according to strict rules
- Apply least privilege
- Keep dependencies updated
			`,
		},
		{
			ID:          "rce--02",
			Title:       "Command Injection",
			Description: "Understand command injection vulnerabilities, how attackers exploit them, and how to prevent them in your applications.",
			Level:       "Beginner",
			Objective:   "Understand how command injection vulnerabilities work.",
			Content: `Command injection is a type of vulnerability that occurs when an application executes OS commands based on user input without proper validation
or sanitization.
Attackers can exploit this vulnerability to execute arbitrary commands on the underlying operating system.

Defensive measures:
- Avoid using user input directly in OS commands
- Use parameterized queries or safe APIs
- Implement proper input validation and sanitization
			`,
		},
		{
			ID:          "rce--03",
			Title:       "Secure Input Handling",
			Description: "Learn how to handle user input securely to prevent RCE and command injection vulnerabilities in your applications.",
			Level:       "Intermediate",
			Objective:   "Learn how to handle user input securely.",
			Content: `Input validation checks whether data matches
the expected format and permitted values.

Good practices:
- Define explicit input constraints
- Reject unexpected input
- Use safe APIs
- Treat validation as one layer of defense`,
		},
	}
}

func FindById(id string) (Lesson, bool) {
	for _, lesson := range GetAll() {
		if lesson.ID == id {
			return lesson, true
		}
	}
	return Lesson{}, false
}

func List() {
	fmt.Println("\n=== RCE Lab | Lessons ===")

	for _, lesson := range GetAll() {
		fmt.Printf("[%s] %s (%s)\n", lesson.ID, lesson.Title, lesson.Level)
		fmt.Printf("    %s\n\n", lesson.Description)
	}
}

func Show(id string) bool {
	lesson, found := FindById(id)
	if !found {
		return false
	}

	fmt.Println("\n=== " + lesson.Title + " ===")
	fmt.Println("ID:", lesson.ID)
	fmt.Println("Level:", lesson.Level)
	fmt.Println("\nObjective:")
	fmt.Println(lesson.Objective)
	fmt.Println("\nLesson:")
	fmt.Println(lesson.Content)

	return true
}
