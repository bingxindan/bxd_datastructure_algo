package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	response := sortedArrayToBST(nums)
	fmt.Println(response)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	// 选择任意一个中间位置数字作为根节点
	mid := (left + right + rand.Intn(2)) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}
