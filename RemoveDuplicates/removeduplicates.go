package main

import (
	"fmt"
	"strings"
)

func main() {
	// var seriesString string
	// fmt.Println("Enter the series of comma separated string")
	// fmt.Scanf("%s", &seriesString)
	// _, err := fmt.Scan(&seriesString)

	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }
	seriesString := "aaa, aaa, bbb,ccc,aaa"
	splitStrings := strings.Split(seriesString, ",")
	// Trim whitespace from each string in the slice
	for i := range splitStrings {
		splitStrings[i] = strings.TrimSpace(splitStrings[i])
	}
	fmt.Println(splitStrings)
	// var arrString []string
	// aaa,aaa,bbb,ccc,aaa
	// arrString = append(arrString, manyStrings...)
	// for _, element := range manyStrings {
	// 	arrString = append(arrString, element)
	// }
	// fmt.Println(arrString)
	// result := removeDuplicates(arrString)
	// fmt.Println(result)
	// // sort.Slice(seriesString)
	// sort.Slice(result, func(i, j int) bool {
	// 	return result[i] < result[j]
	// })
	// fmt.Println(result)

}

// func removeDuplicates(elements []string) []string {
// 	encountered := map[string]bool{} // Map to keep track of encountered elements
// 	result := []string{}             // Slice to store unique elements

// 	for _, v := range elements {
// 		if encountered[v] == false {
// 			encountered[v] = true
// 			result = append(result, v)
// 		}
// 	}

// 	return result
// }
