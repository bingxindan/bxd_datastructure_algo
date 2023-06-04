package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		if fast.Next == nil || slow.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}

	return true
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head.Next

	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
