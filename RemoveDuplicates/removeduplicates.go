package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// Create a scanner to read input from standard input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to enter comma-separated strings
	fmt.Println("Enter comma-separated strings:")

	// Read the input
	if scanner.Scan() {
		// Get the input line and split it by comma
		input := scanner.Text()
		splitStrings := strings.Split(input, ",")

		// Trim whitespace from each string in the slice
		for i, str := range splitStrings {
			splitStrings[i] = strings.TrimSpace(str)
		}

		// Remove duplicates
		uniqueStrings := make(map[string]bool)
		for _, str := range splitStrings {
			uniqueStrings[str] = true
		}

		// Convert the unique strings map to a slice
		var uniqueSlice []string
		for str := range uniqueStrings {
			uniqueSlice = append(uniqueSlice, str)
		}

		// Sort the unique strings
		sort.Strings(uniqueSlice)

		// Print the sorted and unique strings
		fmt.Println("Sorted and unique strings:", uniqueSlice)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
