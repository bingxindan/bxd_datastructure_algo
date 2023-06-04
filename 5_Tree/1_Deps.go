package main

import "fmt"

type Deps struct {
	Val   int
	ParId int
}

type tree struct {
	Val         int
	left, right int
}

func setTree(head map[int]tree, newNode Deps) map[int]tree {
	val := newNode.Val
	pId := newNode.ParId

	if node, ok := head[pId]; ok {
		if node.left == 0 {
			node.left = val
		} else if node.left > val {
			node.right = node.left
			node.left = val
		} else {
			node.right = val
		}
		head[pId] = node
	} else {
		head[pId] = tree{
			Val:   pId,
			left:  val,
			right: 0,
		}
	}

	return head
}

func getTree(trees map[int]tree, root int, data *[]int) {
	node, ok := trees[root]
	if !ok {
		return
	}
	if _, okL := trees[node.left]; !okL {
		*data = append(*data, node.left)
	}
	if _, okR := trees[node.right]; !okR {
		*data = append(*data, node.right)
	}
	getTree(trees, node.left, data)
	getTree(trees, node.right, data)
}

func main() {
	// 3
	// 4,5
	// 6,7,8,9
	a := Deps{Val: 5, ParId: 3}
	b := Deps{Val: 3, ParId: 0}
	c := Deps{Val: 4, ParId: 3}
	d := Deps{Val: 6, ParId: 4}
	e := Deps{Val: 7, ParId: 4}
	f := Deps{Val: 8, ParId: 5}
	g := Deps{Val: 9, ParId: 5}
	h := []Deps{a, b, c, d, e, f, g}
	k := 3

	res := map[int]tree{}
	data := []int{}

	for i := 0; i < len(h); i++ {
		res = setTree(res, h[i])
	}

	getTree(res, k, &data)

	fmt.Println("da:", data, "res:", res)

}
