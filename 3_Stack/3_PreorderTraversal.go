package main

func main() {

}

func preorderTraversal(root *TreeNode) []int {
	var values = make([]int, 0)
	var p1, p2 *TreeNode = root, nil

	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				values = append(values, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			values = append(values, p1.Val)
		}
		p1 = p1.Right
	}

	return values
}
