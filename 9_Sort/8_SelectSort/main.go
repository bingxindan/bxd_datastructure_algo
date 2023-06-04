package main

import "fmt"

func main() {
	nums := []int{39, 2, 5, 23}
	response := selectSort(nums)
	fmt.Println("response: ", response)
}

func selectSort(nums []int) []int {
	n := len(nums)

	if n <= 1 {
		return nums
	}

	for i := 0; i < n-1; i++ {
		minNumIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minNumIndex] {
				minNumIndex = j
			}
		}
		nums[i], nums[minNumIndex] = nums[minNumIndex], nums[i]
	}

	return nums
}
