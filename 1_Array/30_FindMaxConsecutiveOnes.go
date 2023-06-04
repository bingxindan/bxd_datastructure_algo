package main

import "fmt"

func main() {
	nums := []int{1, 0, 1, 1, 0, 1}
	response := findMaxConsecutiveOnes(nums)
	fmt.Println(response)
}

func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	var maxCnt int
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			maxCnt = max(maxCnt, cnt)
			cnt = 0
		}
	}
	maxCnt = max(maxCnt, cnt)
	return maxCnt
}
