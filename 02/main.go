package main

import (
	"fmt"
)

func Print(a int) int {
	if a <= 2 {
		return a
	}

	// return Print(a - 1) + Print(a - 2)
	return a * 300
}

func main() {
	fmt.Println(Print(2))
}