package stringutil

import "fmt"


func ReadString() string {
    var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)
    return input
}