package main

import "fmt"

func main() {
	nums := []int{39, 2, 5, 23}
	response := insertSort(nums)
	fmt.Println("response: ", response)
}

func insertSort(nums []int) []int {
	n := len(nums)

	if n <= 1 {
		return nums
	}

	for i := 1; i < n; i++ {
		num := nums[i]

		j := i - 1
		for ; j >= 0; j-- {
			if num < nums[j] {
				nums[j+1] = nums[j]
			} else {
				break
			}
		}
		nums[j+1] = num
	}

	return nums
}
