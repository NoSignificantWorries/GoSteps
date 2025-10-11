package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 float64

	fmt.Scan(&x1, &y1, &x2, &y2)

	var x float64 = x2 - x1
	var y float64 = y2 - y1

	fmt.Println(math.Sqrt(x*x + y*y))
}
