package main

import (
	"fmt"
	"math"
)

func main() {
	var num int

	fmt.Scan(&num)
	
	num = int(math.Abs(float64(num)))

	n1 := float64(num / 100)
	n2 := float64(num % 100 / 10)
	n3 := float64(num % 10)

	fmt.Println(int(math.Max(n1, math.Max(n2, n3))))
}
