package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Scanf("%d,%d", &num1, &num2)
	// var arr [num1][num2]int
	// var arr1 [num1][num2]int
	var arr1 [][]int
	fmt.Println(arr1)
	var i, j int
	// fmt.Println(arr[0][0])
	for i = 0; i < num2; i++ {
		for j = 0; j < num1; j++ {
			arr1[num2][num1] = i * j
			fmt.Println(arr1)
		}
	}
	fmt.Println(arr1)
}
