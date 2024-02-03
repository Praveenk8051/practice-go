package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var binaryInput string

	fmt.Print("Enter binary numbers separated by commas: ")
	_, err := fmt.Scan(&binaryInput)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Split the input string into individual binary numbers
	binaryNumbers := strings.Split(binaryInput, ",")

	// Convert each binary number to decimal and print the result
	fmt.Println("Decimal equivalents:")
	for _, binaryNumber := range binaryNumbers {
		// Trim any leading or trailing spaces
		binaryNumber = strings.TrimSpace(binaryNumber)

		// Convert binary string to integer
		decimalNumber, err := strconv.ParseInt(binaryNumber, 2, 64)

		if err != nil {
			fmt.Printf("Error converting %s to decimal: %v\n", binaryNumber, err)
			continue
		}

		fmt.Printf("%s: %d\n", binaryNumber, decimalNumber)
	}
}
