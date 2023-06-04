package main

func main() {

}

func kthToLast(head *ListNode, k int) int {
	slow, fast := head, head

	for ; k > 0; k-- {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow.Val
}
