package leetcode

import (
	"container/heap"
	"testing"
)



type PQ []*ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	// 小顶
	return pq[i].Val < pq[j].Val
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}


func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	sentry := &ListNode{}

	tail := sentry

	pq := make(PQ, 0, len(lists))

	for _, v := range lists {
		if v == nil {
			continue
		}
		heap.Push(&pq, v)
	}
	heap.Init(&pq)

	for len(pq) > 0 {
		cur := heap.Pop(&pq).(*ListNode)

		tail.Next = cur
		if cur.Next != nil {
			heap.Push(&pq, cur.Next)
		}
		cur.Next = nil
		tail = cur
	}

	return sentry.Next
}

func TestMerge(t *testing.T) {
	var l1 *ListNode

	//l2 := &ListNode{
	//	Val: 1,
	//	Next: &ListNode{
	//		Val: 3,
	//		Next: &ListNode{
	//			Val: 4,
	//		},
	//	},
	//}
	//
	//l3 := &ListNode{
	//	Val: 2,
	//	Next: &ListNode{
	//		Val: 6,
	//	},
	//}

	lists := []*ListNode{
		l1,
	}
	t.Log(mergeKLists(lists))
}