package main

func main() {

}

func largestValues(root *TreeNode) []int {
	var dfs func(*TreeNode, int)
	var ans = make([]int, 0)
	dfs = func(node *TreeNode, curHeight int) {
		if node == nil {
			return
		}
		if curHeight == len(ans) {
			ans = append(ans, node.Val)
		} else {
			ans[curHeight] = max(ans[curHeight], node.Val)
		}
		dfs(node.Left, curHeight+1)
		dfs(node.Right, curHeight+1)
	}
	dfs(root, 0)
	return ans
}
