package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	D := math.Pow(b, 2) - 4*a*c

	if D < 0 {
		return
	} else if D == 0 {
		x := -b / (2 * a)
		fmt.Println(x)
	} else {
		D = math.Sqrt(D)
		x1 := (-b - D) / (2 * a)
		x2 := (-b + D) / (2 * a)

		fmt.Printf("%g\n%g\n", math.Min(x1, x2), math.Max(x1, x2))
	}
}
