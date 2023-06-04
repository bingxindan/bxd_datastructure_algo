package main

import "fmt"

func main() {
	response := generateParenthesis(1)
	fmt.Println(response)
}

var response []string

func generateParenthesis(n int) []string {
	generate(0, 0, n, "")
	return response
}

func generate(left, right, n int, s string) {
	if left == n && right == n {
		response = append(response, s)
		return
	}

	if left < n {
		generate(left+1, right, n, s+"(")
	}
	if left > right {
		generate(left, right+1, n, s+")")
	}
}
