package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	printMenu()
}
func printMenu() {
	fmt.Println("Enter any comma seperated numbers")
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("Unknown Value")
	} else {
		fmt.Printf("Entered numbers are %s", input)
		num := sliceFromStrings(input)
		print(num)
	}
}
func sliceFromStrings(input string) []int {

	numbers := strings.Split(input, ",")

	length := len(numbers)
	var num = make([]int, length)

	for index, v := range numbers {
		s := strings.Trim(v, " ")
		num[index], _ = strconv.Atoi(s)
	}
	return num
}
