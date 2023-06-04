package main

import (
	"container/heap"
	"sort"
)

func main() {

}

var factors = []int{2, 3, 5}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}

	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}
