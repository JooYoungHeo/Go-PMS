package main

import "fmt"

func Move(from int, to int) {
	fmt.Printf("%d -> %d\n", from, to)
}

func Hanoi(n int, from int, pass int, to int) {
	if n == 1 {
		Move(from, to)
	} else {
		Hanoi(n-1, from, to, pass)
		Move(from, to)
		Hanoi(n-1, pass, from, to)
	}
}

func main() {
	n := 3
	Hanoi(n, 1, 3, 2)
}
