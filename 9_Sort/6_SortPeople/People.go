package main

import (
	"fmt"
	"sort"
)

func main() {
	names := []string{"Mary", "John", "Emma"}
	heights := []int{180, 165, 170}
	response := sortPeople(names, heights)
	fmt.Printf("response: %+v\n", response)
}

func sortPeople(names []string, heights []int) []string {
	length := len(names)
	mark := make(map[int]string, length)

	for i := 0; i < length; i++ {
		mark[heights[i]] = names[i]
	}

	sort.Ints(heights)

	for i := length - 1; i >= 0; i-- {
		names[length-1-i] = mark[heights[i]]
	}

	return names
}
