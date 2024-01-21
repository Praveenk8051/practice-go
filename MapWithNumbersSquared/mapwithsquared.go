package main

import (
	"fmt"
)

func main() {
	printMenu()
}
func printMenu() {
	fmt.Println("Enter any numer")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Unknown Value")
	} else {
		fmt.Printf("%v", MapNumbers(num))
	}
}

// Ex003 returns a map with numbers are their squared values
func MapNumbers(num int) map[int]int {
	// create a map with the size of n
	var numbers = make(map[int]int, num)

	for i := 1; i <= num; i++ {
		numbers[i] = i * i
	}
	return numbers
}
