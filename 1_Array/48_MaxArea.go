package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	//height := []int{1, 1}
	//response := maxArea1(height)
	response := maxArea2(height)
	fmt.Println(response)
}

func maxArea2(height []int) int {
	n := len(height)
	left, right := 0, n-1

	var response int

	for left < right {
		var tmp int
		if height[left] < height[right] {
			tmp = height[left] * (right - left)
			left++
		} else {
			tmp = height[right] * (right - left)
			right--
		}
		if response < tmp {
			response = tmp
		}
	}

	return response
}

func maxArea1(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	left, right := 0, n-1
	var (
		sum  int
		area int
	)

	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}

		if sum < area {
			sum = area
		}
	}

	return sum
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res, area := 0, 0

	for left < right {
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if res < area {
			res = area
		}
	}

	return res
}
