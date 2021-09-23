package main

import "fmt"

func fibonacci(n int) int {
	return fibonacciNumberTailHelper(0, 1, n)
}

func fibonacciNumberTailHelper(fn1, fn2, n int) int {
	if n < 1 {
		return fn1
	}
	return fibonacciNumberTailHelper(fn2, fn1+fn2, n-1)
}

func main() {
	fmt.Println(fibonacci(9))
	fmt.Println(fibonacci(10))
	fmt.Println(fibonacci(19))
}
