package main

import "fmt"

func main() {
	var a, n float64

	fmt.Scan(&a, &n)

	fmt.Println(a * (1 - n/100.0))
}
