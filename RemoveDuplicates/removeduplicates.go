package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("Enter any comma seperated words")
	var input string
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("Unknown Value")
	} else {
		output := removeDuplicates(input)
		fmt.Print(output)
	}
}
func removeDuplicates(input string) string {
	fmt.Println(input)
	// make a map and split the input string
	m := make(map[string]bool)
	temp := strings.Split(input, " ")

	// add values to map as keys (making them unique)
	for _, v := range temp {
		m[v] = true
	}

	// convert map back into slice
	var result []string
	for key := range m {
		result = append(result, key)
	}

	// sort slice and join back toghether
	sort.Strings(result)
	output := strings.Join(result, " ")

	return output
}
