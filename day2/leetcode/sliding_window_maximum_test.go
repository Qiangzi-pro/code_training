package leetcode

import (
	"container/heap"
	"testing"
)

type Item struct {
	Value    int // 元素的值，可以是任意字符串。
	Index int         // 元素在队列中的优先级。
}

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 大顶
	return pq[j].Value < pq[i].Value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func heapSolution(nums []int, k int) []int {
	q := make(PriorityQueue, 0, 100)
	var res = make([]int, 0, len(nums)-k+1)
	for idx, v := range nums[:k-1] {
		q = append(q, &Item{
			Value: v,
			Index: idx,
		})
	}

	heap.Init(&q)

	max := &Item{
		Index: -1,
	}
	for idx := k-1; idx < len(nums); idx++{
		heap.Push(&q, &Item{
			Value: nums[idx],
			Index: idx,
		})
		for max.Index < idx-k+1 || nums[idx] > max.Value {
			max = heap.Pop(&q).(*Item)
		}
		res = append(res, max.Value)
	}

	return res
}

func dequeSolution(nums []int, k int) []int {
	deque := make([]int, 0, len(nums))
	res := make([]int, 0, len(nums)-k+1)

	for i, v := range nums {
		for len(deque) != 0 && nums[deque[len(deque)-1]] < v {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		if deque[0] == i-k {
			deque = deque[1:]
		}

		if i >= k-1 {
			res = append(res, nums[deque[0]])
		}
	}
	return res
}

func maxSlidingWindow(nums []int, k int) []int {
	return dequeSolution(nums, k)
}

func TestSliding(t *testing.T) {
	t.Log(maxSlidingWindow([]int{1,3,-1,-3,5,3,6,7}, 3))
}