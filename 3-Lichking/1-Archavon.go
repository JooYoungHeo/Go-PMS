package main

import "fmt"

var (
	start = rune(44032)
	end   = rune(55204)
)

func HasConsonantSuffix(s string) bool {
	numEnds := 28
	result := false

	for _, r := range s {
		if start <= r && r < end {
			index := int(r - start)
			result = index%numEnds != 0
		}
	}

	return result
}

func main() {
	fruits := []string{"사과", "바나나", "토마토", "수박"}

	for _, fruit := range fruits {
		post := "는"

		if HasConsonantSuffix(fruit) {
			post = "은"
		}

		fmt.Printf("%s%s 맛있다.\n", fruit, post)
	}
}
