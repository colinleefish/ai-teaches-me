// Week 1 Quiz: Go Basics, Variables, and Types
// Interactive quiz following CS61A assessment style

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	ID       int
	Question string
	Options  []string
	Answer   int
	Points   int
	Category string
}

func main() {
	fmt.Println("🎯 Week 1 Quiz: Go Basics, Variables, and Types")
	fmt.Println("=" + strings.Repeat("=", 50))
	
	quiz := createWeek1Quiz()
	score := runQuiz(quiz)
	
	fmt.Printf("\n📊 Final Score: %d/%d (%.1f%%)\n", score, getTotalPoints(quiz), float64(score)/float64(getTotalPoints(quiz))*100)
	
	if float64(score)/float64(getTotalPoints(quiz)) >= 0.8 {
		fmt.Println("🎉 Excellent! You're ready for Week 2!")
	} else {
		fmt.Println("📚 Review the material and try again. Target: 80%+")
	}
}

func createWeek1Quiz() []Question {
	return []Question{
		{
			ID:       1,
			Question: "Which of the following is the correct way to declare a variable in Go?",
			Options: []string{
				"var name string = \"Alice\"",
				"name := \"Alice\"",
				"var name = \"Alice\"",
				"All of the above",
			},
			Answer:   4,
			Points:   5,
			Category: "Variables",
		},
		{
			ID:       2,
			Question: "What is the zero value of an int in Go?",
			Options: []string{
				"nil",
				"0",
				"undefined",
				"\"\"",
			},
			Answer:   2,
			Points:   5,
			Category: "Types",
		},
		{
			ID:       3,
			Question: "Which statement about Go's type system is correct?",
			Options: []string{
				"Go has dynamic typing like Python",
				"Go has static typing checked at compile time",
				"Go variables can change types at runtime",
				"Go doesn't have a type system",
			},
			Answer:   2,
			Points:   10,
			Category: "Types",
		},
		{
			ID:       4,
			Question: "What happens when you run: go run main.go?",
			Options: []string{
				"Creates an executable file",
				"Compiles and runs the program temporarily",
				"Only checks syntax",
				"Formats the code",
			},
			Answer:   2,
			Points:   5,
			Category: "Basics",
		},
		{
			ID:       5,
			Question: "Which of these is NOT a valid Go basic type?",
			Options: []string{
				"int",
				"float64",
				"string",
				"list",
			},
			Answer:   4,
			Points:   5,
			Category: "Types",
		},
		{
			ID:       6,
			Question: "What does the := operator do in Go?",
			Options: []string{
				"Assigns a value to an existing variable",
				"Declares and initializes a new variable",
				"Compares two values",
				"Both A and B",
			},
			Answer:   2,
			Points:   10,
			Category: "Variables",
		},
		{
			ID:       7,
			Question: "In Go, what is the difference between an array and a slice?",
			Options: []string{
				"No difference",
				"Arrays have fixed size, slices are dynamic",
				"Arrays are dynamic, slices have fixed size",
				"Arrays are for strings, slices are for numbers",
			},
			Answer:   2,
			Points:   10,
			Category: "Types",
		},
		{
			ID:       8,
			Question: "What is the correct way to create a map in Go?",
			Options: []string{
				"map[string]int{\"key\": 1}",
				"make(map[string]int)",
				"var m map[string]int = make(map[string]int)",
				"All of the above",
			},
			Answer:   4,
			Points:   10,
			Category: "Types",
		},
		{
			ID:       9,
			Question: "Which command formats Go code according to Go standards?",
			Options: []string{
				"go format",
				"go fmt",
				"gofmt",
				"go style",
			},
			Answer:   3,
			Points:   5,
			Category: "Basics",
		},
		{
			ID:       10,
			Question: "What is the scope of a variable declared with := inside an if block?",
			Options: []string{
				"Global scope",
				"Function scope",
				"Block scope only",
				"Package scope",
			},
			Answer:   3,
			Points:   10,
			Category: "Variables",
		},
		{
			ID:       11,
			Question: "How do you convert a string to an integer in Go?",
			Options: []string{
				"int(\"123\")",
				"strconv.Atoi(\"123\")",
				"string.ToInt(\"123\")",
				"parse(\"123\")",
			},
			Answer:   2,
			Points:   10,
			Category: "Types",
		},
		{
			ID:       12,
			Question: "What is the purpose of the go.mod file?",
			Options: []string{
				"Contains main function",
				"Defines module and manages dependencies",
				"Stores configuration settings",
				"Contains test cases",
			},
			Answer:   2,
			Points:   10,
			Category: "Basics",
		},
		{
			ID:       13,
			Question: "In Go, which of these will cause a compile error?",
			Options: []string{
				"var x int = 10",
				"x := 10",
				"var x := 10",
				"var x int; x = 10",
			},
			Answer:   3,
			Points:   10,
			Category: "Variables",
		},
		{
			ID:       14,
			Question: "What is the zero value of a boolean in Go?",
			Options: []string{
				"true",
				"false",
				"0",
				"nil",
			},
			Answer:   2,
			Points:   5,
			Category: "Types",
		},
		{
			ID:       15,
			Question: "How do you declare a constant in Go?",
			Options: []string{
				"const pi = 3.14",
				"constant pi = 3.14",
				"final pi = 3.14",
				"let pi = 3.14",
			},
			Answer:   1,
			Points:   5,
			Category: "Variables",
		},
	}
}

func runQuiz(questions []Question) int {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	
	for i, q := range questions {
		fmt.Printf("\n📝 Question %d/%d (%s - %d points):\n", i+1, len(questions), q.Category, q.Points)
		fmt.Println(q.Question)
		fmt.Println()
		
		for j, option := range q.Options {
			fmt.Printf("  %d. %s\n", j+1, option)
		}
		
		fmt.Print("\nYour answer (1-4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		answer, err := strconv.Atoi(input)
		if err != nil || answer < 1 || answer > 4 {
			fmt.Println("❌ Invalid input. No points awarded.")
			continue
		}
		
		if answer == q.Answer {
			fmt.Printf("✅ Correct! (+%d points)\n", q.Points)
			score += q.Points
		} else {
			fmt.Printf("❌ Incorrect. Correct answer: %d. %s\n", q.Answer, q.Options[q.Answer-1])
		}
	}
	
	return score
}

func getTotalPoints(questions []Question) int {
	total := 0
	for _, q := range questions {
		total += q.Points
	}
	return total
}

// 🎯 QUIZ CATEGORIES:
// - Basics: Go installation, commands, modules
// - Variables: Declaration, scope, constants
// - Types: Basic types, conversions, zero values

// 📊 SCORING:
// - 90-100%: Outstanding mastery
// - 80-89%: Good understanding, ready to proceed
// - 70-79%: Needs some review
// - Below 70%: Significant review needed

// 🚀 AFTER THE QUIZ:
// - Review any missed concepts
// - Practice with the exercise files
// - Retake if score is below 80%
// - Proceed to Week 2 if ready 