package main

import (
	"fmt"
)

type CQueue struct {
	// 栈 A，用于添加元素
	stackA []interface{}
	// 栈 B，用于取出元素
	stackB []interface{}
}

// CQueue 的构造函数
func ConstructorII() CQueue {
	// 返回一个新的 CQueue
	return CQueue{}
}

// 往队列插入整数
func (this *CQueue) AppendTail(value interface{}) {
	// 在 stackA 的末尾追加 value
	this.stackA = append(this.stackA, value)
}

// 从队列取出整数并删除
func (this *CQueue) DeleteHead() interface{} {
	// 如果 stackB 没有元素则从 stackA 中取出所有
	if len(this.stackB) == 0 {
		// 如果 stackA 里也没有元素，则队列爲空返回 -1
		if len(this.stackA) == 0 {
			return -1
		}
		// 将 stackA 的所有元素转移到 stackB
		for len(this.stackA) > 0 {
			// 获取 stackA 最末尾元素的下标
			index := len(this.stackA) - 1
			// 获取 stackA 最末尾元素的值 value
			value := this.stackA[index]
			// 向 stackB 的末尾追加 value
			this.stackB = append(this.stackB, value)
			// 从 stackA 中裁剪出末尾元素
			this.stackA = this.stackA[:index]
		}
	}
	// 这时候表示 stackB 内已有元素
	// 获取 stackB 最末尾元素的下标
	index := len(this.stackB) - 1
	// 获取 stackB 最末尾元素的值 value
	value := this.stackB[index]
	// 从 stackB 中裁剪出末尾元素
	this.stackB = this.stackB[:index]
	// 返回 value
	return value
}

func main() {
	obj := ConstructorII()

	//var ele = []int{2, 23, 3453, 564, 5, 67, 4, 123, 66, 57, 34, 809, 8765, 4243, 5765}
	var ele = []interface{}{"", "", 5, 2, "", ""}

	for _, v := range ele {
		obj.AppendTail(v)
	}

	fmt.Println(11, obj.stackA)

	for {
		param := obj.DeleteHead()

		fmt.Println(22, param)

		if param == -1 {
			break
		}
	}
}
