package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	node *TreeNode
	prt  *TreeNode
}

func convertBiNode(root *TreeNode) *TreeNode {
	var tree = &Tree{}
	tree.node = &TreeNode{}
	tree.TreeTo(root)
	return tree.node.Right
}

func (t *Tree) TreeTo(next *TreeNode) {
	if next == nil {
		return
	}

	t.TreeTo(next.Left)

	if t.prt == nil {
		t.prt = next
		t.node.Right = next
	} else {
		t.prt.Right = next
		t.prt = next
	}

	next.Left = nil
	t.TreeTo(next.Right)
}
