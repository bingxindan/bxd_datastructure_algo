package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	i := len(matrix) - 1
	j := 0

	for i > -1 {
		if j < len(matrix[i]) {
			if target < matrix[i][j] {
				i--
			} else if target > matrix[i][j] {
				j++
			} else if target == matrix[i][j] {
				return true
			}
		} else {
			return false
		}
	}

	return false
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30}}
	res := findNumberIn2DArray(matrix, 9)
	fmt.Println("r:", res)
}
