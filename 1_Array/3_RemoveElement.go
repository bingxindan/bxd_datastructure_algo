package main

import "fmt"

func main() {
	var nums = []int{1, 2, 2, 3, 0, 4, 2}
	val := 2
	//response := removeElement(nums, val)
	response := removeElementII(nums, val)
	fmt.Println(response)
}

func removeElement(nums []int, val int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	var slow = 0

	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[slow] = nums[i]
			slow++
		}
	}

	return slow
}

func removeElementII(nums []int, val int) int {
	left, right := 0, len(nums)

	if right == 0 {
		return 0
	}

	for left < right {
		if nums[left] == val {
			nums[left] = nums[right-1]
			right--
		} else {
			left++
		}
	}

	return left
}
