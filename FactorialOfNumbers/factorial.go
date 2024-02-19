package main

import (
	"fmt"
)

func main() {
	var num int
	num = printNum(num)
	fmt.Printf("You have entered %d\n", num)
	result := factorial(num)
	printResult(result, num)
}
func printNum(num int) int {
	fmt.Println("Enter the number below 20, to find factorial")

	_, err := fmt.Scanf("%d", &num)
	checkifGreater(num, err)
	return num
}

func checkifGreater(num int, err error) {
	if err != nil {
		fmt.Println("Unknown value")
	}
	if !(num < 20) {
		printNum(num)
	}
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
