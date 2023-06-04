package main

func main() {

}

func numIslands(grid [][]byte) int {
	var ans int
	m, n := len(grid), len(grid[0])

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n {
			return
		}
		if grid[r][c] != '1' {
			return
		}
		grid[r][c] = '2'
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == '1' {
				ans++
				dfs(r, c)
			}
		}
	}

	return ans
}
