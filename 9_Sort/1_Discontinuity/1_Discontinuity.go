package main

import "fmt"

func main() {
	a := []int{3, 1, 2, 0, 5}

	// 排序
	for i := 0; i < len(a); i++ {

		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	if a[0] != 0 {
		fmt.Println(0)
		return
	}
	if a[len(a)-1] != len(a) {
		fmt.Println(len(a))
		return
	}

	for m := 1; m < len(a); m++ {
		if a[m]-a[m-1] != 1 {
			fmt.Println(a[m] - 1)
			return
		}
	}

	fmt.Println(a)
}
