package hangul

import "fmt"

func ExampleHasCopnsonantSuffix() {
	fmt.Println(HasConsonantSuffix("Go 언어"))
	// Output:
	// false
}
