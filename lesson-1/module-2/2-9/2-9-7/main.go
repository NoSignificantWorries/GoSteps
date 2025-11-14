package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	fmt.Scan(&a, &b)

	var c float64 = math.Sqrt(a*a + b*b)

	fmt.Println(a + b + c)
}
