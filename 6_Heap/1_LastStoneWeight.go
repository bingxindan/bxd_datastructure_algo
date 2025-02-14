package main

import (
	"container/heap"
)

func main() {

}

func lastStoneWeight(stones []int) int {
	q := &hp{stones}
	heap.Init(q)

	for q.Len() > 1 {
		x, y := q.pop(), q.pop()
		if x > y {
			q.push(x - y)
		}
	}

	if q.Len() > 0 {
		return q.IntSlice[0]
	}

	return 0
}

func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }
