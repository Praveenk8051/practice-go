package main

import (
	"bufio"
	"os"
)

func ReadString() string {
	var input string = ""
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	temputils.PrintString(input)
}
