package main

import (
	"fmt"
	"strings"
)

func main() {
	// Input
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	// Make it case insensitive
	lowerInput := strings.ToLower(input)

	// Check
	startsWithI := strings.HasPrefix(lowerInput, "i")
	endsWithN := strings.HasSuffix(lowerInput, "n")
	containsA := strings.Contains(lowerInput, "a")

	// Print result based on the conditions
	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
