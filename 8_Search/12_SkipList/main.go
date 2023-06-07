package main

import (
	"fmt"
	"math/rand"
)

func main() {
	obj := Constructor()
	obj.Add(3)
	obj.Add(4)
	obj.Add(5)
	obj.Add(6)
	obj.Add(7)
	return
	response1 := obj.Search(1)
	fmt.Println("res1: ", response1)
	obj.Add(3)
	obj.Add(4)
	obj.Add(5)
	obj.Add(6)
	obj.Add(7)
	response2 := obj.Search(4)
	fmt.Println("res2: ", response2)
}

const (
	maxLevel = 32
	pFactor  = 0.25
)

type SkiplistNode struct {
	val     int
	forward []*SkiplistNode
}

type Skiplist struct {
	head  *SkiplistNode
	level int
}

func Constructor() Skiplist {
	return Skiplist{&SkiplistNode{-1, make([]*SkiplistNode, maxLevel)}, 0}
}

func (Skiplist) randomLevel() int {
	lv := 1
	for lv < maxLevel && rand.Float64() < pFactor {
		lv++
	}
	return lv
}

func (this *Skiplist) Search(target int) bool {
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 找到第i层小于且最接近target的元素
		for curr.forward[i] != nil && curr.forward[i].val < target {
			curr = curr.forward[i]
		}
	}
	curr = curr.forward[0]
	// 检测当前元素的值是否等于target
	return curr != nil && curr.val == target
}

func (this *Skiplist) Add(num int) {
	update := make([]*SkiplistNode, maxLevel)
	for i := range update {
		update[i] = this.head
	}

	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		// 找到第i层小于且最接近num的元素
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	lv := this.randomLevel()
	this.level = max(this.level, lv)
	newNode := &SkiplistNode{num, make([]*SkiplistNode, lv)}
	for i, node := range update[:lv] {
		// 对第i层的状态进行更新，将当前元素的forward指向新的节点
		newNode.forward[i] = node.forward[i]
		node.forward[i] = newNode
	}
}

func (this *Skiplist) Erase(num int) bool {
	update := make([]*SkiplistNode, maxLevel)
	curr := this.head
	for i := this.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].val < num {
			curr = curr.forward[i]
		}
		update[i] = curr
	}
	curr = curr.forward[0]

	// 如果值不存在，则返回false
	if curr == nil || curr.val != num {
		return false
	}

	for i := 0; i < this.level && update[i].forward[i] == curr; i++ {
		// 对第i层的状态进行更新，将forward指向被删除节点的下一跳
		update[i].forward[i] = curr.forward[i]
	}
	// 更新当前的level
	for this.level > 1 && this.head.forward[this.level-1] == nil {
		this.level--
	}

	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
