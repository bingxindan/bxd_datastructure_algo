package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	response := sortedSquares(nums)
	fmt.Println(response)
}

func sortedSquares1(nums []int) []int {
	length := len(nums)
	response := make([]int, length)
	i, j := 0, length-1

	for pos := length - 1; pos >= 0; pos-- {
		v, w := nums[i]*nums[i], nums[j]*nums[j]
		if v > w {
			response[pos] = v
			i++
		} else {
			response[pos] = w
			j--
		}
	}

	return response
}

func sortedSquares(nums []int) []int {
	length := len(nums)
	response := make([]int, 0)

	if length == 0 {
		return response
	}

	for i := 0; i < length; i++ {
		num := nums[i]
		response = append(response, num*num)
	}

	sort.Ints(response)

	return response
}
