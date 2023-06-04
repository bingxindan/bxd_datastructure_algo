package main

import "fmt"

type Student struct {
	Name string
	Age  int
	Info []string
}

func main() {
	/*var f1 float64 = 0.3
	var f2 float64 = 0.3

	if f1 == f2 {
		fmt.Println(1)
	} else {
		fmt.Println(2)
	}*/

	a := Student{
		Name: "minping",
		Age:  30,
	}
	b := Student{
		Name: "minping",
		Age:  30,
	}
	fmt.Println(a, b)
}
