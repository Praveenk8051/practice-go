package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a scanner to read input from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to enter lines of text
	fmt.Println("Enter lines of text (press Enter twice to end):")

	var lines []string

	// Read lines of text until an empty line is entered
	for scanner.Scan() {
		line := scanner.Text()

		// If an empty line is entered, break the loop
		if line == "" {
			break
		}

		// Add the line to the slice
		lines = append(lines, line)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
		return
	}

	// Convert each line to uppercase and print
	for _, line := range lines {
		fmt.Println(strings.ToUpper(line))
	}
}
