package main

import (
	"fmt"
	"math"
)

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)
	fmt.Println("Entered Value", input)
	var a [3]int
	var Q [3]int
	fmt.Sscanf(input, "%d,%d,%d", &a[0], &a[1], &a[2])
	fmt.Println("String converted values are: ", a[0], a[1], a[2])
	var C = 50
	var H = 30
	for val, num := range a {
		fmt.Println(num, val)
		Q[val] = int(math.Sqrt(float64((2 * C * num) / H)))
	}
	fmt.Println("FinalAnswer", Q)

}
