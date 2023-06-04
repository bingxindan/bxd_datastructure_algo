package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	response := containsDuplicate(nums)
	fmt.Println(response)
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]struct{}, 0)

	for _, v := range nums {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}

	return false
}
