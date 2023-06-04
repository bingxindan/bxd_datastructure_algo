package main

import "fmt"

func main() {
	nums := []int{1, 1}
	response := findDisappearedNumbers(nums)
	fmt.Println(response)
}

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	ans := make([]int, 0)
	for _, v := range nums {
		v = (v - 1) % n
		nums[v] += n
	}
	for i, v1 := range nums {
		if v1 <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
