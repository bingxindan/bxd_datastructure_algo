package main

import (
	"fmt"
)

func SafeClose(ch chan int) (closed bool) {
	defer func() {
		if recover() != nil {
			closed = false
			fmt.Println("2")
		}
	}()
	// 如果 ch 是一个已经关闭的，会 panic 的，然后被 recover 捕捉到；
	fmt.Println("1")
	close(ch)
	return true
}

func fooChan1() {
	a := make(chan string, 3)
	a <- "a"
	a <- "b"
	a <- "c"
	close(a)

	fmt.Println(111, <-a)
	fmt.Println(222, <-a)
	fmt.Println(333, <-a)
}

func fooChan2() {
	a := make(chan string)
	a <- "a"
	fmt.Println(111, <-a)
}

func fooChan3() {
	chs := make(chan int, 3)
	chs <- 1
	chs <- 2
	chs <- 3

	//chs <- 4
}

func main() {
	fooChan3()

	//fooChan2()

	/*ch := make(chan int, 1)

	go func() {
		fmt.Println("wait 3s...")
		time.Sleep(3 * time.Second)
		ch <- 1
	}()

	data := <-ch

	fmt.Println(data)*/

	/*wg := sync.WaitGroup{}
	c := make(chan int)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	wg.Add(4)

	for i := range c {
		wg.Done()
		fmt.Println(i)
	}

	close(c)
	wg.Wait()*/

	/*out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			out <- rand.Intn(5)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println(i)
		}
	}()
	wg.Wait()*/

	/*c := make(chan int, 4)
	c <- 1
	c <- 2
	c <- 3
	c <- 4

	d := make(chan int, 3)
	d <- 5
	d <- 6
	d <- 7

	SafeClose(c)
	close(d)
	SafeClose(c)

	isC := false
	isD := false

	for i := 0; i < 20; i++ {
		if i == 0 {
			isC = true
		}

		if isC {
			vC, okC := <-c
			if vC > 0 && okC {
				fmt.Println("c:", vC)
				isD = true
				isC = false
			}
		}

		if isD {
			vD, okD := <-d
			if vD > 0 && okD {
				fmt.Println("d:", vD)
				isC = true
				isD = false
			}
		}
	}*/

	/*a := make(chan int, 4)
	a <- 1
	a <- 2
	a <- 3
	a <- 4

	b := make(chan int, 6)
	b <- 5
	b <- 6
	b <- 7
	b <- 8
	b <- 9
	b <- 10

	close(a)
	close(b)

	for i := 0; i < 12; i++ {
		select {
		case v, ok := <-a:
			fmt.Println("v:", v, "ok:", ok)
		case v1, ok1 := <-b:
			fmt.Println("v1:", v1, "ok1:", ok1)
		default:
			fmt.Println("end")
			break
		}
	}*/

	/*for v := range a {
		fmt.Println("v:", v)
	}*/

	/*leng := len(a)
	fmt.Println("a:", leng)*/
}
