package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	//response := runningSum(nums)
	response := runningSum1(nums)
	fmt.Println(response)
}

func runningSum1(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

func runningSum(nums []int) []int {
	response := make([]int, 0)
	length := len(nums)

	if length == 0 {
		return response
	}

	for i := 0; i < length; i++ {
		var sum int
		for j := 0; j <= i; j++ {
			num1 := nums[j]
			sum += num1
		}
		response = append(response, sum)
	}

	return response
}
