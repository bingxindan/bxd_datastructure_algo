package main

import "fmt"

func main() {
	obj := ConstructorI(3)
	obj.Put(1, 1)
	obj.Put(2, 2)
	response1 := obj.Get(1)
	fmt.Println(response1)
	obj.Put(3, 3)
	obj.Put(4, 4)
	response2 := obj.Get(4)
	fmt.Println(response2)
	response3 := obj.Get(1)
	fmt.Println(response3, obj.cache)

}

type LRUCacheI struct {
	size     int // 长度
	capacity int // 容量
	cache    map[int]*DLinkedNode
	head     *DLinkedNode
	tail     *DLinkedNode
}

type DLinkedNode struct {
	key   int
	value int
	prev  *DLinkedNode
	next  *DLinkedNode
}

func ConstructorI(capacity int) LRUCacheI {
	l := LRUCacheI{
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{key: 0, value: 0},
		tail:     &DLinkedNode{key: 0, value: 0},
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCacheI) Put(key, value int) {
	// 校验是否存在
	val, ok := this.cache[key]
	if ok {
		// 如果存在，则和get类似，将节点更新到链表头
		val.value = value
		this.moveToHead(val)
		return
	}

	// 如果不存在，则创建新节点，先加入到链表头，在加入到hash表中
	node := &DLinkedNode{key: key, value: value}
	this.cache[key] = node
	this.addToHead(node)
	this.size++

	// 如果超过了，则要删除链表尾值
	if this.size > this.capacity {
		// 删除尾部节点
		removed := this.removeTail()
		// 删除hash中的值
		delete(this.cache, removed.key)
		// 当前容量size-1
		this.size--
	}
}

func (this *LRUCacheI) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

func (this *LRUCacheI) Get(key int) int {
	// 判断key是否存在
	val, ok := this.cache[key]
	if !ok {
		return -1
	}

	// 存在，先把值入到链表头部
	this.moveToHead(val)

	// 在将值返回
	return val.value
}

// 将节点移动到表头
func (this *LRUCacheI) moveToHead(node *DLinkedNode) {
	// 先移除节点
	this.removeNode(node)
	// 再添加到节点头
	this.addToHead(node)
}

// 移除节点
func (this *LRUCacheI) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 添加头节点
func (this *LRUCacheI) addToHead(node *DLinkedNode) {
	// 更新节点前指针
	node.prev = this.head
	// 更新节点后指针
	node.next = this.head.next
	// 更新原来虚拟头节点的后指针节点的前指针
	this.head.next.prev = node
	// 更新原来虚拟头节点的后指针
	this.head.next = node
}
