package main

import "fmt"

func main() {
	fmt.Println("Enter the number between 2000-3200 you want to check!")
	var num int
	_, err := fmt.Scanf("%d", &num)

	if err != nil {
		fmt.Println("Unknown Value")
	}
	if !(num >= 2000 && num <= 3200) {
		fmt.Println("Please enter the number between 2000-3200")
	}
	if !(divisibleBy7(num) && notDivisibleBy5(num)) {
		fmt.Println("Number doesn't satify the needs")
	} else {
		fmt.Println("The number is Divisible By 7 and Not Divisible By 5")
	}

}

func divisibleBy7(num int) bool {
	return (num%7 == 0)
}
func notDivisibleBy5(num int) bool {
	return num%5 != 0
}
