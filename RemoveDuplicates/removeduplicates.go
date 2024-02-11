package main

import (
	"fmt"
	"sort"
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
	arrString = append(arrString, manyStrings...)
	// for _, element := range manyStrings {
	// 	arrString = append(arrString, element)
	// }
	// fmt.Println(arrString)
	result := removeDuplicates(arrString)
	fmt.Println(result)
	// sort.Slice(seriesString)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})
	fmt.Println(result)

	// fmt.Println("Sorted alphanumerically:", seriesString)
	// for _, manyString := range manyStrings {

	// }
}
func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{} // Map to keep track of encountered elements
	result := []string{}             // Slice to store unique elements

	for _, v := range elements {
		if encountered[v] == false {
			encountered[v] = true
			result = append(result, v)
		}
	}

	return result
}
