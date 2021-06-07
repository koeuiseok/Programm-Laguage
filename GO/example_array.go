package main

import "fmt"

func main() {
	array := [5]string{
		"array0",
		"array1",
		"array2",
		"array3",
		"array4",
	}
	fmt.Println(array)
	a := array[1:4]
	b := array[0:2]
	array[1] = "change"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(array)
}
