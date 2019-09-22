package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	for i := 0; i < 20; i++ {
		delta := (z*z - x) / (2 * z)
		if math.Abs(delta) < 1e-8 {
			break
		}
		fmt.Printf("i: %2d - delta: %0.20f\n", i, delta)
		z -= delta
	}
	return z
}

func main() {
	x := math.Pow(334, 2)
	y := Sqrt(float64(x))
	fmt.Printf("y: %10.0f, y**2: %10.0f\n", y, math.Pow(y, 2))

}
