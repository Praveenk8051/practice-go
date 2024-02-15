package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter comma-separated strings:")

	if scanner.Scan() {
		input := scanner.Text()

		uniqueStrings := make(map[string]bool)

		for _, val := range input {
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
