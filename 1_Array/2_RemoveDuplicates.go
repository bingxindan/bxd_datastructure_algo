package main

import "fmt"

func main() {
	var nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	response := removeDuplicates(nums)
	fmt.Println(response)
}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	var slow = 1

	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}
