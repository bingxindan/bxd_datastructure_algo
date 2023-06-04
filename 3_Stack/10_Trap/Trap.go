package main

import "fmt"

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	response := trap(height)
	fmt.Println(response)
}

func trapI(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	leftM, rightM := 0, 0
	var ans int

	for left < right {
		leftM = max(leftM, height[left])
		rightM = max(rightM, height[right])

		if height[left] < height[right] {
			ans += leftM - height[left]
			left++
		} else {
			ans += rightM - height[right]
			right--
		}
	}

	return ans
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	var ans int

	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
