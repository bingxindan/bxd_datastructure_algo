package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 3, 2, 1}
	response := maopao(nums)
	fmt.Println(response)
}

func maopao(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
