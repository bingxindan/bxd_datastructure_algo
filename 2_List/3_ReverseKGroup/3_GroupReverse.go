package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	list := createListI(arr)

	head := reverseKGroupI(list, 4)
	for head != nil {
		fmt.Println("node: ", head.Val)
		head = head.Next
	}

	/*newHead := ReverseKGroup1(tail, 3)
	ShowGroup(newHead)*/
}

func reverseKGroupI(head *ListNode, k int) *ListNode {
	// 定义虚表头
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	// 循环head，找出k个元素
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}
		// 记录下一个节点
		tmp := tail.Next
		// 反转链表
		head, tail = reverseList(head, tail)
		prev.Next = head
		tail.Next = tmp
		prev = tail
		head = tail.Next
	}

	return dummyHead.Next
}

func reverseList(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head

	for prev != tail {
		tmp := p.Next
		p.Next = prev
		prev = p
		p = tmp
	}

	return tail, head
}

func createListI(arr []int) *ListNode {
	n := len(arr)

	if n == 0 {
		return &ListNode{}
	}

	var head *ListNode

	for i := 0; i < n; i++ {
		node := &ListNode{Val: arr[i]}
		node.Next = head
		head = node
	}

	return head
}

func ReverseKGroup1(head *ListNode, k int) *ListNode {
	var entry = &ListNode{0, head}
	last := entry
	for head != nil {
		// 求组
		tail := GroupGet(head, k)
		if tail == nil {
			break
		}
		nextHead := tail.Next

		// 翻转
		GroupReverse(head, nextHead)

		// 连接
		last.Next = tail
		head.Next = nextHead
		last = head

		// 下一组
		head = nextHead
	}

	return entry.Next
}

func ReverseKGroup(head *ListNode, k int) *ListNode {
	var entry = &ListNode{0, head}
	last := entry
	for head != nil {
		tail := GroupGet(head, k)
		if tail == nil {
			break
		}
		nextHead := tail.Next

		GroupReverse(head, nextHead)

		last.Next = tail
		head.Next = nextHead
		last = head

		head = nextHead
	}
	return entry.Next
}

func GroupReverse(head, stop *ListNode) *ListNode {
	last := head
	for head != stop {
		node := head.Next
		head.Next = last
		last = head
		head = node
	}
	return last
}

func GroupGet(head *ListNode, k int) *ListNode {
	for head != nil {
		k--
		if k == 0 {
			return head
		}
		head = head.Next
	}
	return nil
}

func ShowGroup(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
