package main

func main() {

}

func subsets(nums []int) [][]int {
	set := make([]int, 0)
	ans := make([][]int, 0)
	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), set...))
			return
		}
		set = append(set, nums[cur])
		dfs(cur + 1)
		set = set[:len(set)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return ans
}
