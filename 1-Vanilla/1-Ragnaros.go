package main

import "fmt"

func sing(n int) {
	repetition := 1

	for repetition <= n {
		pants := repetition * 10
		sword := (repetition + 1) * 10

		fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", pants, sword)
		repetition++
	}
}

func main() {
	sing(3)
}
