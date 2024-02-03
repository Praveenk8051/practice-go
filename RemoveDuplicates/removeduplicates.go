package main

import (
	"fmt"
	"sort"
)

func main() {
	var seriesString []string
	fmt.Scanf("Enter the series of string", &seriesString)
	_, err := fmt.Scan(&seriesString)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Split the input string into individual binary numbers
	// manyStrings := strings.Split(seriesString, ",")
	// fmt.Println(manyStrings)
	// sort.Slice(seriesString)
	sort.Slice(seriesString, func(i, j int) bool {
		return seriesString[i] < seriesString[j]
	})
	fmt.Println("Sorted alphanumerically:", seriesString)
	// for _, manyString := range manyStrings {

	// }
}
