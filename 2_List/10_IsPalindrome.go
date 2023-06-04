package main

func main() {

}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	p1 := head
	p2 := secondHalfStart
	result := true

	for result && p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	firstHalfEnd.Next = reverseList(secondHalfStart)

	return result
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
