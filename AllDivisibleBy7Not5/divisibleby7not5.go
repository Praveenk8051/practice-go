package main

import "fmt"

func main() {
	printMenu()
}
func printMenu() {
	fmt.Println("Enter any range of Numbers divisible by 7 and Not divisible by 5")
	var num1, num2 int
	_, err := fmt.Scanf("%d, %d", &num1, &num2)
	fmt.Println(err)
	if err != nil {
		fmt.Println("Unknown Value")
	} else {
		fmt.Printf("Entered numbers are %d, %d\n", num1, num2)
		checkNums(num1, num2)
	}
}

func checkNums(num1, num2 int) {
	if num1 <= num2 {
		fmt.Println("Enter num1 higher than num2")
		printMenu()
	}
	printNums(num1, num2)

}

func printNums(num1, num2 int) {
	var i int
	for i = num1; i <= num2; i++ {
		if !(divisibleBy7(i) && notDivisibleBy5(i)) {
			continue
		} else {
			fmt.Printf("%d ", i)
		}
	}
}

func divisibleBy7(num int) bool {
	return (num%7 == 0)
}
func notDivisibleBy5(num int) bool {
	return num%5 != 0
}
