package main

import "fmt"

func main() {

}

func numWays(n int, relation [][]int, k int) int {
	edges := make([][]int, n)
	for _, r := range relation {
		fmt.Println(edges, r)
	}
	return 0
}
