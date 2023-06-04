package main

import "fmt"

func main() {
	//s := "()"
	//response := isValid(s)
	//response := isValidI(s)

	/*response := make([]string, 0)
	response = append(response, "a")
	response = append(response, "b")
	response = append(response, "c")
	response = append(response, "d")
	response = append(response, "e")

	fmt.Println(response, response[0], len(response), response[:len(response)-1], response[2:])*/
}

func isValidI(s string) bool {
	length := len(s)
	if length&1 == 1 {
		return false
	}

	inputMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, 0)

	for i := 0; i < length; i++ {
		str := s[i]
		if inputMap[str] > 0 {
			if len(stack) == 0 || inputMap[str] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str)
		}
	}

	return len(stack) == 0
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack []byte

	for i := 0; i < n; i++ {
		fmt.Println("aaaaaaa", s[i], pairs[s[i]], stack, pairs)
		if pairs[s[i]] > 0 {
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
