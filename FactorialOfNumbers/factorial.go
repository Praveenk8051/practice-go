package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter the number to find factorial")
	var num int
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Unknown value")
	}
	fmt.Printf("You have entered %d\n", num)
	result := factorial(num)
	printResult(result, num)
}

func factorial(num int) int {
	var temp = 1
	var i int
	for i = num; i >= 1; i-- {
		temp = temp * i
		num = i
	}
	return temp
}

func printResult(result, num int) {
	fmt.Printf("The factorial of %d = %d\n", num, result)
}
