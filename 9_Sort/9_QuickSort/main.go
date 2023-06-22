package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{39, 2, 5, 23}
	response := quickSort(nums, 0, len(nums)-1)
	fmt.Println("response: ", response)
}

func sortArray(nums []int) []int {
	quick(&nums, 0, len(nums)-1)
	return nums
}

func quick(arr *([]int), i, j int) {
	if i >= j {
		return
	}
	mid := partition(arr, i, j)
	quick(arr, i, mid-1)
	quick(arr, mid+1, j)
}

func partition(arr *([]int), i, j int) int {
	// 随机选取支点
	p := rand.Intn(j-i+1) + i
	nums := *arr
	nums[i], nums[p] = nums[p], nums[i]
	for i < j {
		// 修改原来的 nums[j] >= nums[i]，增加交换频率
		for nums[i] < nums[j] && i < j {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		// 修改原来的 nums[j] >= nums[i]，增加交换频率
		for nums[i] < nums[j] && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return i
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
