package main

import (
	"fmt"
	"strconv"
)

func main() {
	nums := []int{0, 1, 2, 4, 5, 7}
	response := summaryRanges(nums)
	fmt.Println(response)
}

func summaryRanges(nums []int) []string {
	var ans = make([]string, 0)
	for i, n := 0, len(nums); i < n; {
		left := i
		for i++; i < n && nums[i-1]+1 == nums[i]; i++ {

		}
		s := strconv.Itoa(nums[left])
		if left < i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		ans = append(ans, s)
	}
	return ans
}
