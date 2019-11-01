package main

import (
	"fmt"
)

var (
	myres = make(map[int]int, 20)
	ch = make(chan int)
)

func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	ch <- res
}

func main() {
	for i := 1; i <= 20; i++ {
		go factorial(i)
		myres[i]= <- ch
	}
	close(ch)
	for i, v := range myres {
		fmt.Printf("myres[%d] = %d\n", i, v)
	}
}
