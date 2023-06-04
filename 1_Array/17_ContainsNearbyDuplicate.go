package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	response := containsNearbyDuplicate(nums, k)
	fmt.Println(response)
}

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]struct{}, 0)

	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}

	return false
}
