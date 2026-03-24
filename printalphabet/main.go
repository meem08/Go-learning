package main

import "fmt"

// Write a program that prints the Latin alphabet in lowercase on a single line.
// A line is a sequence of characters preceding the end of line character ('\n').

func main() {
	for i := 0; i <= 9; i++ {
		fmt.Print(i)
	}
	fmt.Println()
	// IsNegative(-3)
	fmt.Println(IsNegative(56))
}

func IsNegative(n int) bool {
	if n < 0 {
		return true
	}
	return false
}
