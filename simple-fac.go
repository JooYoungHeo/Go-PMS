package main

import "fmt"

func facItr(n int) int {
	result := 1
	for n > 0 {
		result *= n
		fmt.Printf("n: %d, result: %d\n", n, result)
		n--
	}
	return result
}

func main() {
	fmt.Println(facItr(5))
}
