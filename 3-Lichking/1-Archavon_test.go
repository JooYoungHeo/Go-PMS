package archavon

import "fmt"

func ExampleModifyHangulPost() {
	fruits := []string{"사과", "바나나", "딸기", "수박"}
	for _, fruit := range fruits {
		post := ModifyHangulPost(fruit)
		fmt.Printf("%s%s 맛있다.\n", fruit, post)
	}
	// Output:
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 딸기는 맛있다.
	// 수박은 맛있다.
}
