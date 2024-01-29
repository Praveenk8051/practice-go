package main

import (
	"classwith2methodsinterface/stringutil"
	"fmt"
)

func main() {

	userInput := stringutil.ReadString()

	upperCaseString := stringutil.PrintString(userInput)

	fmt.Println("Original Input:", userInput)
	fmt.Println("UpperCase Output:", upperCaseString)
}
