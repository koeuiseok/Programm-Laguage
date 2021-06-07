package main

import (
	"fmt"
)

var cnt int = 0

func increment(c chan int) {
	for i := 0; i < 100000; i++ {
		cnt = cnt + 1;
	}
	c <- cnt
}



func main() {
	c := make(chan int, 1)
	
	go increment(c)
	cnt = <- c
	increment(c)
	cnt = <- c
	fmt.Println(cnt)
	
}
