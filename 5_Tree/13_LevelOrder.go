package main

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	q := []*TreeNode{root}
	for i := 0; len(q) > 0; i++ {
		ret = append(ret, []int{})
		p := make([]*TreeNode, 0)
		for j := 0; j < len(q); j++ {
			node := q[j]
			ret[i] = append(ret[i], node.Val)
			if node.Left != nil {
				p = append(p, node.Left)
			}
			if node.Right != nil {
				p = append(p, node.Right)
			}
		}
		q = p
	}

	return ret
}
