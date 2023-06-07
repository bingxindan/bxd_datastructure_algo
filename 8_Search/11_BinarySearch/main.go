package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	response := binarySearch(nums, 7)
	fmt.Println("response: ", response)
}

func binarySearch(nums []int, value int) int {
	n := len(nums)

	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == value {
			return mid
		} else if nums[mid] > value {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
