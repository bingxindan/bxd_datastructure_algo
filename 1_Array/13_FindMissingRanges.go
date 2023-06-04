package main

import "fmt"

func main() {
	nums := []int{-1}
	lower := -1
	upper := -1
	response := findMissingRanges(nums, lower, upper)
	fmt.Println(response)
}

func findMissingRanges(nums []int, lower, upper int) [][]int {
	n := len(nums)
	m := n + 2
	tmpNums := make([]int, m)
	tmpNums[0] = lower - 1
	tmpNums[m-1] = upper + 1
	copy(tmpNums[1:m-1], nums)
	var ret [][]int

	for i := 1; i < m; i++ {
		if tmpNums[i-1] != tmpNums[i] {
			r := getRange(tmpNums[i-1], tmpNums[i])
			if len(r) > 0 {
				ret = append(ret, r)
			}
		}
	}

	return ret
}

func getRange(lower, upper int) []int {
	r := upper - lower
	if r < 2 {
		return []int{}
	} else if r == 2 {
		return []int{lower + 1, lower + 1}
	} else {
		return []int{lower + 1, upper - 1}
	}
}
