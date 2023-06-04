package main

func main() {

}

func kthLargest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode)
	var res = -1

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		k--
		if k == 0 {
			res = node.Val
			return
		}
		dfs(node.Left)
	}

	dfs(root)

	return res
}
