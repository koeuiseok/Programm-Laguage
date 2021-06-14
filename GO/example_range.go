package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i, _ := range pow { // 0,1,2,... 배열의 순서만
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow { // 1,2,4,8,,,  배열의 값만
		fmt.Printf("%d\n", value)
	}
}
