package main

import (
	"fmt"
	"strings"
)

func main() {
	var seriesString string
	fmt.Scanf("Enter the series of comma separated string %s", &seriesString)
	_, err := fmt.Scan(&seriesString)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Split the input string into individual binary numbers
	manyStrings := strings.Split(seriesString, ",")
	var arrString []string
	for _, element := range manyStrings {
		arrString = append(arrString, element)
	}
	fmt.Println(arrString)
	// sort.Slice(seriesString)
	// sort.Slice(seriesString, func(i, j int) bool {
	// 	return seriesString[i] < seriesString[j]
	// })
	// fmt.Println("Sorted alphanumerically:", seriesString)
	// for _, manyString := range manyStrings {

	// }
}
