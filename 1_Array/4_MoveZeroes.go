package main

import "fmt"

func main() {
	nums := []int{0}
	response := moveZeroes1(nums)
	fmt.Println(response)
}

func moveZeroes1(nums []int) []int {
	left, right, n := 0, 0, len(nums)
	response := make([]int, 0)

	if n == 0 {
		return response
	}

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	return nums
}

func moveZeroes(nums []int) []int {
	n := len(nums)

	if n == 0 {
		return nil
	}

	left, right := 0, 0

	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	return nums
}
