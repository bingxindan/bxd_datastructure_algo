package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 2
	response := searchInsert(nums, target)
	fmt.Println(response)
}

func searchInsert(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	ans := n

	for left <= right {
		mid := (right-left)>>1 + left
		if target <= nums[mid] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return ans
}
