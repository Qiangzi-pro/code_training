package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	var p1, p2 *ListNode = head, head

	for p2.Next != nil && p2.Next.Next != nil {
		p2 = p2.Next.Next
		p1 = p1.Next

		if p1 == p2 {
			return true
		}
	}
	return false
}
