package main

import "fmt"

func depth(grid *[][]int, i, j int) int {
	if i > 0 && i < len(*grid) && j > 0 && j < len((*grid)[i]) && (*grid)[i][j] == 1 {
		(*grid)[i][j] = 0
		num := 1 + depth(grid, i-1, j) + depth(grid, i+1, j) + depth(grid, i, j-1) + depth(grid, i, j+1)
		return num
	}
	return 0
}

func main() {
	grid := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}

	grid = [][]int{{1}}

	if len(grid) <= 1 && len(grid[0]) <= 1 {
		//return grid[0][0]
	}

	res := 0

	for i, v1 := range grid {
		for j, v2 := range v1 {
			if v2 == 1 {
				num := depth(&grid, i, j)
				if num > res {
					res = num
				}
			}
		}
	}

	fmt.Println("r:", res)
}
