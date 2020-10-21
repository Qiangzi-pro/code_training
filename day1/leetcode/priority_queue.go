package leetcode

import (
	"container/heap"
	"fmt"
)

// Item 是优先队列中包含的元素。
type Item struct {
	Value    interface{} // 元素的值，可以是任意字符串。
	Priority int         // 元素在队列中的优先级。
	// 元素的索引可以用于更新操作，它由 heap.Interface 定义的方法维护。
	Index int // 元素在堆中的索引。
}

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 小顶
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}

// 更新函数会修改队列中指定元素的优先级以及值。
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

// 这个示例首先会创建一个优先队列，并在队列中包含一些元素
// 接着将一个新元素添加到队列里面，并对其进行操作
// 最后按优先级有序地移除队列中的各个元素。
func main() {
	// 一些元素以及它们的优先级。
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}

	// 创建一个优先队列，并将上述元素放入到队列里面，
	// 然后对队列进行初始化以满足优先队列（堆）的不变性。
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item{
			Value:    value,
			Priority: priority,
			Index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 插入一个新元素，然后修改它的优先级。
	item := &Item{
		Value:    "orange",
		Priority: 1,
	}
	heap.Push(&pq, item)
	//pq.update(item, item.Value, 5)

	// 以降序形式取出并打印队列中的所有元素。
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s ", item.Priority, item.Value)
	}
}