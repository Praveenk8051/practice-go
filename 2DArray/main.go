package main

import "fmt"

func main() {
	var arr = [5][2]int{{1, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
	var arr1 [5][2]int
	var i, j int
	fmt.Println(arr[3][0])
	for i = 0; i < 2; i++ {
		for j = 0; j < 5; j++ {
			arr1[j][i] = i * j
		}
	}
	fmt.Println(arr1)
}
