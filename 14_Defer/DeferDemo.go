package main

import "fmt"

func main() {
	a := DeferHandle()
	fmt.Println(a)
}

func DeferHandle() int {
	defer func() {
		fmt.Println(111)
	}()

	fmt.Println(222)

	defer func() {
		fmt.Println(333)
		defer func() {
			fmt.Println(444)
		}()
	}()

	return 1
}
