package main

import "fmt"

func Fibonacci(n int) int {
	p, q := 0, 1

	if n == 0 {
		return p
	}
	if n == 1 {
		return q
	}

	for i := 2; i < n; i++ {
		p, q = q, p+q
	}

	return q
}

func main() {
	n := 11
	result := Fibonacci(n)

	fmt.Println(result)
}
