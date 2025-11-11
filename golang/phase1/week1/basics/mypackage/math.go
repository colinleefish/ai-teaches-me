package mypackage

import "fmt"

func Add(a, b int) int {
	fmt.Println("Adding two numbers")
	return a + b
}

func Subtract(a, b int) int {
	fmt.Println("Subtracting two numbers")
	return a - b
}
