package main

import (
	"fmt"
)

var cnt int = 0

func increment() {
	for i := 0; i < 1000000; i++ {
		cnt = cnt + 1;
	}
}

func decrement() {
	for i := 0; i < 1000000; i++ {
		cnt = cnt - 1;
	}
}


func main() {
	go increment()
	increment()
	fmt.Println(cnt)
	
	go decrement()
	decrement()
	fmt.Println(cnt)
}
