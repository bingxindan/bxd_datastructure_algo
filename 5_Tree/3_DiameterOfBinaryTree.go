package main

func main() {

}

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(node *TreeNode) int
	var ans int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return -1
		}
		lLen := dfs(node.Left) + 1
		rLen := dfs(node.Right) + 1
		ans = max(ans, lLen+rLen)
		return max(lLen, rLen)
	}
	dfs(root)
	return ans
}
