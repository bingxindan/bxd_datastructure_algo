package main

import "fmt"

func main() {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	response := intersection(nums1, nums2)
	fmt.Println(response)
}

func intersection(nums1, nums2 []int) []int {
	set1 := make(map[int]struct{}, 0)
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := make(map[int]struct{}, 0)
	for _, v1 := range nums2 {
		set2[v1] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}

	intersections := make([]int, 0)
	for v2 := range set1 {
		if _, has := set2[v2]; has {
			intersections = append(intersections, v2)
		}
	}

	return intersections
}

func intersect(nums1, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersect(nums2, nums1)
	}

	m := make(map[int]int, 0)
	for _, num := range nums1 {
		m[num]++
	}

	intersections := make([]int, 0)
	for _, num1 := range nums2 {
		if m[num1] > 0 {
			intersections = append(intersections, num1)
			m[num1]--
		}
	}

	return intersections
}
