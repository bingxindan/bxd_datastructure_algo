package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 0, -2, -1, 1, 2}
	response := threeSum(nums)
	fmt.Println(response)
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}

/*func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			if second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return ans
}*/
