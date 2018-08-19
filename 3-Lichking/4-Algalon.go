package main

import "fmt"

func put(list []string, target string) []string {
	list = append(list, target)
	return list
}

func get(list []string) []string {
	list = list[1:]
	return list
}

func main() {
	list := []string{"a", "b"}
	list = put(list, "c")
	list = get(list)

	fmt.Println(list)
}
