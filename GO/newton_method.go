package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var min float64 = 1 << 10;
	var z float64 =1;
  
	for cnt:=0; cnt < 10; cnt++ {
		z -= (math.Pow(z, 2) - x)/(2*z)
		z = math.Abs(z)
		if z < min {
			min = z
		}
	}
	return z
}

 

func main() {
	fmt.Println(Sqrt(77))
	fmt.Println(math.Sqrt(77))
}
