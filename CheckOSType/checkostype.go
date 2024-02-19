package main

import (
	"fmt"
	"math"
	"runtime"
)

func main() {
	if runtime.GOARCH == "amd64" || runtime.GOARCH == "386" {
		fmt.Println("The system is running in 32-bit mode.")
		minInt32 := math.MinInt32
		maxInt32 := math.MaxInt32

		fmt.Printf("Range of int32: %d to %d\n", minInt32, maxInt32)
	} else if runtime.GOARCH == "amd64" || runtime.GOARCH == "amd64" {
		minInt64 := math.MinInt64
		maxInt64 := math.MaxInt64

		fmt.Printf("Range of int32: %d to %d\n", minInt64, maxInt64)
		fmt.Println("The system is running in 64-bit mode.")
	} else {
		fmt.Println("Unknown architecture:", runtime.GOARCH)
	}
}
