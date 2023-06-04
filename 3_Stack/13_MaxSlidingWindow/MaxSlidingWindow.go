package main

import "fmt"

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	response := maxSlidingWindow(nums, 3)
	fmt.Println(response)
}

func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	var deque []int
	ans := make([]int, 0, n-k+1)

	for i, num := range nums {
		// 将出界的下标删除
		if i >= k && deque[0] <= i-k {
			deque = deque[1:]
		}

		// 将 i 添加到队列，并保证队列是递减的。
		for len(deque) > 0 && nums[deque[len(deque)-1]] <= num {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if i >= k-1 {
			ans = append(ans, nums[deque[0]])
		}
	}

	return ans
}
