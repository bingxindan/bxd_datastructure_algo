package main

import "fmt"

func main() {
	ret := minDistance("horse", "ros")
	fmt.Println(ret)
}

func minDistance(word1, word2 string) int {
	m := len(word2)
	f := make([]int, m+1)
	for j := 1; j <= m; j++ {
		f[j] = j
	}
	for _, x := range word1 {
		pre := f[0]
		f[0]++
		for j, y := range word2 {
			if x == y {
				f[j+1], pre = pre, f[j+1]
			} else {
				f[j+1], pre = min(min(f[j+1], f[j]), pre)+1, f[j+1]
			}
		}
	}
	return f[m]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
