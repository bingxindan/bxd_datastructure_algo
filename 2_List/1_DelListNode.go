package main

import "fmt"

type DelListNode struct {
	Val  int
	Next *DelListNode
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	var head *DelListNode
	for i := 0; i < len(a); i++ {
		node := &DelListNode{
			Val: a[i],
		}
		node.Next = head
		head = node
	}

	last := DelReverse(head)

	list := RemoveLinkListNode(last, 3)

	DelShowReverse(list)
}

// 写一个函数，删除一个未知长度的单向链表倒数第 N 个节点元素。
func RemoveLinkListNode(list *DelListNode, n int) *DelListNode {
	var node = &DelListNode{}
	dummy := list
	i := 1
	for list != nil {
		if i == n {
			// 执行删除
			node.Next = list.Next
			list.Next = nil
			break
		}
		node = list
		list = list.Next
		i++
	}
	return dummy
}

func DelReverse(head *DelListNode) *DelListNode {
	var last *DelListNode
	for head != nil {
		node := head.Next
		head.Next = last
		last = head
		head = node
	}
	return last
}

func DelShowReverse(head *DelListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
