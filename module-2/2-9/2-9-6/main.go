package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Scan(&num)

	fmt.Println(num / math.Pow(2, 13))
}
