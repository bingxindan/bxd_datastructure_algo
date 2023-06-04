package main

import (
	"fmt"
)

// findPairs 返回 nums 里所有差为 k 的唯一组合数。
//
// k 大于 0。
func findPairs(nums []int, k int) int {
	leng := len(nums)
	if leng == 0 {
		return 0
	}
	m := 0
	tmps := make(map[[2]int]int, 0)

	for i := 0; i < leng; i++ {
		for j := i + 1; j < leng; j++ {
			diff := nums[i] - nums[j]
			if diff == k || diff == -k {
				key := [2]int{nums[i], nums[j]}
				if nums[i] > nums[j] {
					key = [2]int{nums[j], nums[i]}
				}
				if _, ok := tmps[key]; ok {
					continue
				}
				tmps[key] = 1

				m++
			}
		}
	}

	fmt.Println("nums", nums, "m", m)

	return 0
}

func main() {
	var cases = []struct {
		nums   []int
		k      int
		expect int
	}{
		{[]int{}, 2, 0},
		{[]int{1}, 2, 0},
		{[]int{1, 3}, 2, 1},          // (1, 3)
		{[]int{3, 1, 4, 1, 5}, 2, 2}, // (1, 3), (3, 5)
		{[]int{1, 1, 1, 1, 1}, 2, 0},
		{[]int{1, 1, 1, 2, 2}, 1, 1}, // (1, 2)
	}
	for _, c := range cases {
		if r := findPairs(c.nums, c.k); r != c.expect {
			//fmt.Printf("failed: case=%+v autual=%d\n", c, r)
		}
	}
}
