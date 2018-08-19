package kelthuzad

import "fmt"

func ExampleCheckIndex() {
	list := []string{"java", "c++", "c#", "nodejs", "perl", "ruby", "go", "c"}
	target := "nodejs"
	idx := CheckIndex(list, target)

	fmt.Println(idx)

	// Output
	// 3
}
