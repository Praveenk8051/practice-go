package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter comma-separated strings:")

	if scanner.Scan() {
		input := scanner.Text()

		splitStrings := strings.Split(input, ",")

		// Trim whitespace from each string in the slice
		for i, str := range splitStrings {
			splitStrings[i] = strings.TrimSpace(str)
		}
		uniqueStrings := make(map[string]bool)

		for val := range input {
			uniqueStrings[val] = true
		}
		var uniqueSlice []string
		for str := range uniqueStrings {
			uniqueSlice = append(uniqueSlice, str)
		}
		sort.Strings(uniqueSlice)

		fmt.Println("Sorted and unique strings:", uniqueSlice)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading standard input:", err)
	}
}
