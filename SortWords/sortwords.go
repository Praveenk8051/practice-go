package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Println("Enter comma-separated strings:")

	if scanner.Scan(){
		input := scanner.Text()
		uniqueStrings := make(map[string]bool)

		for _ str := range input{
			uniqueStrings[str] = true
		}
		
}
