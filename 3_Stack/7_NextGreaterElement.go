package main

func main() {

}

func nextGreaterElement(nums1, nums2 []int) []int {
	mp := make(map[int]int, 0)
	stack := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}

	res := make([]int, len(nums1))
	for i, num := range nums1 {
		res[i] = mp[num]
	}

	return res
}
