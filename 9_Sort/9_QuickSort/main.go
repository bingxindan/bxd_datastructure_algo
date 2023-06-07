package main

import "fmt"

func main() {
	nums := []int{39, 2, 5, 23}
	response := quickSort(nums, 0, len(nums)-1)
	fmt.Println("response: ", response)
}

func quickSort(nums []int, left, right int) []int {
	if left < right {
		ptIndex := partitionWay(nums, left, right)
		quickSort(nums, left, ptIndex-1)
		quickSort(nums, ptIndex+1, right)
	}

	return nums
}

func partitionWay(nums []int, left, right int) int {
	// 先分区，最后把基准换到边界上
	privot := left
	index := privot + 1

	for i := index; i <= right; i++ {
		// 当前值小于基准就交换，大于的不用管
		if nums[privot] > nums[i] {
			nums[index], nums[i] = nums[i], nums[index]
			// 交换后的下一个
			index++
		}
	}

	// nums[index]是大于基准的
	nums[privot], nums[index-1] = nums[index-1], nums[privot]

	return index - 1
}
