package main

import (
	"fmt"
	"sort"
)

func main() {
	g := []int{1, 2}
	s := []int{1, 2, 3}
	response := findContentChildren(g, s)
	fmt.Println(response)
}

func findContentChildren(g, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	m, n := len(g), len(s)
	var ans int

	for i, j := 0, 0; i < m && j < n; i++ {
		for j < n && g[i] > s[j] {
			j++
		}
		if j < n {
			ans++
			j++
		}
	}

	return ans
}
