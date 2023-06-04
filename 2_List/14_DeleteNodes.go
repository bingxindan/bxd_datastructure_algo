package main

func main() {

}

func deleteNodes(head *ListNode, m, n int) *ListNode {
	pre := &ListNode{}
	p := pre
	pre.Next = head

	for pre != nil {
		cnt := 0
		for cnt < m && pre != nil {
			pre = pre.Next
			cnt++
		}
		if pre == nil {
			break
		}
		cnt = 0
		t := pre.Next
		for cnt < n && t != nil {
			t = t.Next
			cnt++
		}
		pre.Next = t
	}

	return p.Next
}
