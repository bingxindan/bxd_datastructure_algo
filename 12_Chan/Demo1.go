package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		chA = make(chan int, 20)
	)

	for i := 0; i < 10; i++ {
		chA <- i
	}

	close(chA)

	for v := range chA {
		fmt.Println(v)
	}

	time.Sleep(4 * time.Second)
}
