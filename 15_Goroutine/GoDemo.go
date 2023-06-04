package main

import (
	"fmt"
	"time"
)

func main() {
	a := 0

	for i := 0; i < 1000; i++ {
		go func() {
			a++
		}()
		time.Sleep(2 * time.Millisecond)
	}

	fmt.Println("a:", a)
}
