package main

import "fmt"

func main() {
	nums := []int{39, 2, 5, 23, 10}
	response := mergeSort(nums)
	fmt.Println("response: ", response)
}

func mergeSort(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	mid := n / 2
	left := nums[:mid]
	right := nums[mid:]

	return merge(mergeSort(left), mergeSort(right))
}

func merge(left, right []int) []int {
	var res []int

	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			res = append(res, left[0])
			// 将头一个直接切出去
			left = left[1:]
		} else {
			res = append(res, right[0])
			right = right[1:]
		}
	}

	if len(left) == 0 {
		res = append(res, right...)
	}
	if len(right) == 0 {
		res = append(res, left...)
	}

	return res
}
