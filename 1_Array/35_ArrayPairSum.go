package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 4, 3, 2}
	response := arrayPairSum(nums)
	fmt.Println(response)
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	var ans int

	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}

	return ans
}
