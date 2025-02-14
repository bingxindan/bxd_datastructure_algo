package main

import "fmt"

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	response := islandPerimeter(grid)
	fmt.Println(response)
}

type pair struct {
	x, y int
}

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func islandPerimeter(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	var ans int
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			for _, d := range dir4 {
				if x, y := i+d.x, j+d.y; x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
					ans++
				}
			}
		}
	}

	return ans
}
