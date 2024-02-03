package main

import "fmt"

func generateArray(X, Y int) [][]int {
	array := make([][]int, X)

	for i := 0; i < X; i++ {
		array[i] = make([]int, Y)
		for j := 0; j < Y; j++ {
			array[i][j] = i * j
		}
	}

	return array
}

func main() {
	var X, Y int

	fmt.Print("Enter the value of X: ")
	fmt.Scan(&X)
	fmt.Print("Enter the value of Y: ")
	fmt.Scan(&Y)

	resultArray := generateArray(X, Y)

	fmt.Println("Generated 2D Array:")
	for _, row := range resultArray {
		fmt.Println(row)
	}
}
