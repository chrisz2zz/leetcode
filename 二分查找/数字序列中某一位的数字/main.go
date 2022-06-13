package main

import "fmt"

func main() {
	fmt.Println(findNthDigit(100))
}

func findNthDigit(n int) int {
	count := 1
	digits := 0

	for n > digits {
		n -= digits
		digits = count * 9
		count++
	}

	fmt.Println(n)
	return 0
}