package main

import (
	"fmt"
	"math"
)

func recursiveDigitsAdded(digits int) int {
	return recursiveDigitsAddedHelper(digits, 0)
}

func recursiveDigitsAddedHelper(digits, lastNumber int) int {
	x := splitAndAdd(digits, 0)
	lastNumber += x
	if x < 10 {
		return lastNumber
	} else {
		return recursiveDigitsAddedHelper(x, lastNumber)
	}
}

func splitAndAdd(digits, total int) int {
	rightNumber := digits % 10
	if digits < 10 {
		total += digits
		return total
	} else {
		digits = int(math.Floor(float64(digits / 10)))
		return splitAndAdd(digits, total+rightNumber)
	}
}

func main() {
	fmt.Println(recursiveDigitsAdded(99999999999884))
}
