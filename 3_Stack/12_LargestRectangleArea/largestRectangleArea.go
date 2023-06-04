package main

import "fmt"

func main() {
	heights := []int{2, 1, 5, 6, 2, 3}
	//heights := []int{0, 9}
	response := largestRectangleArea(heights)
	fmt.Println(response)
}

func largestRectangleAreaI(heights []int) int {
	n := len(heights)
	var (
		stack [][]int
		ans   = 0
	)

	for i := 0; i <= n; i++ {
		height := 0
		if i < n {
			height = heights[i]
		}

		if len(stack) == 0 {
			stack = append(stack, []int{1, height})
			continue
		}

		if height == stack[len(stack)-1][1] {
			stack[len(stack)-1][0] = stack[len(stack)-1][0] + 1
			continue
		}
		if height > stack[len(stack)-1][1] {
			stack = append(stack, []int{1, height})
			continue
		}

		width := 0
		for len(stack) > 0 && height < stack[len(stack)-1][1] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			width += top[0]
			ans = max(ans, width*top[1])
		}
		stack = append(stack, []int{width + 1, height})
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	length := len(heights)
	if length == 0 {
		return 0
	}

	var area int

	for i := 0; i < length; i++ {
		left, right := i, i

		for left-1 >= 0 && heights[left-1] >= heights[i] {
			left--
		}
		for right < length-1 && heights[right+1] >= heights[i] {
			right++
		}

		width := right - left + 1
		areaTmp := width * heights[i]
		if areaTmp > area {
			area = areaTmp
		}
	}

	return area
}
