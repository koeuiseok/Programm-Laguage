package main

import "fmt"

func add() func(int) int {
	out := 0
	return func(x int) int {
		out += x
		return out
	}
}

func main() {
	pos, neg := add(), add()
	for i := 0; i < 15; i++ {
		fmt.Println(
			pos(i),
			neg(-*i),
		)
	}
	fmt.Println(add())  // ??????
	fmt.Println(pos(15))
}
