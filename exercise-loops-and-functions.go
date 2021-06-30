// https://tour.golang.org/flowcontrol/8
package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	oldz := z
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		if z == oldz {
			break
		}
		oldz = z
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

