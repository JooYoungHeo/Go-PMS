package malygos

import "fmt"

func ExampleSortingSlice() {
	testList := []int{10, 5, 3, 4, 8, 9, 1, 2, 7, 6}

	sortedList := SortingSlice(testList)

	fmt.Println(sortedList)

	// Output:
	// [1 2 3 4 5 6 7 8 9 10]
}
