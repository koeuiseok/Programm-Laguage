package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64, a float64) float64 {
	return fn(a, 3)
}

func comp(fn func( float64) float64, inp float64) float64 {
	return fn(inp)
}

func fibonacci(a float64) float64 {
	if a<=1{
		return 1;
	}
	return fibonacci(a-2) + fibonacci(a-1)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot, 3))
	fmt.Println(compute(math.Pow, 3))
	fmt.Println(comp(fibonacci, 5))
}
